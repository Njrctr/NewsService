package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/rpcgen/v2/golang"
	middleware "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"net/http"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"news-service/internal/rest"
	myrpc "news-service/internal/rpc"
	"os"
	"time"
)

type App struct {
	Cfg  *Config
	DB   db.DB
	rest *rest.Handler
	rpc  zenrpc.Server
}

func New(cfg *Config, dbconn db.DB) *App {
	repository := db.NewNewsRepo(dbconn)
	services := newsportal.NewManager(repository)
	handlers := rest.New(services)

	r := myrpc.Init(services)

	return &App{
		Cfg:  cfg,
		DB:   dbconn,
		rest: handlers,
		rpc:  r,
	}
}

type Config struct {
	HttpPort int
	Env      string

	DB *pg.Options
}

func (a *App) Run(ctx context.Context) error {
	elog := log.New(os.Stderr, "E", log.LstdFlags|log.Lshortfile)
	dlog := log.New(os.Stderr, "D", log.LstdFlags|log.Lshortfile)

	srv := a.rest.Init()
	allowDebug := func(param string) middleware.AllowDebugFunc {
		return func(req *http.Request) bool {
			return req.FormValue(param) == "true"
		}
	}

	a.rpc.Use(

		middleware.WithDevel(true),
		middleware.WithAPILogger(dlog.Printf, middleware.DefaultServerName),
		middleware.WithTiming(true, allowDebug("d")),
		middleware.WithSQLLogger(a.DB, true, allowDebug("d"), allowDebug("s")),
		middleware.WithErrorLogger(elog.Printf, middleware.DefaultServerName),
	)

	srv.Any("/rpc", echo.WrapHandler(a.rpc))
	srv.Any("/rpc/doc", echo.WrapHandler(http.HandlerFunc(zenrpc.SMDBoxHandler)))

	gen := rpcgen.FromSMD(a.rpc.SMD())

	srv.GET("/rpc/client.go", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.GoClient(golang.Settings{})))))
	srv.GET("/rpc/client.ts", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.TSClient(nil)))))
	srv.GET("/rpc/open", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.OpenRPC("", "")))))

	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(stopCtx); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("--- Доступные эндпоинты ---")
	for _, route := range srv.Routes() {
		fmt.Printf("Метод: %-7s Путь: %s\n", route.Method, route.Path)
	}
	fmt.Println("---------------------------")

	return srv.Start(fmt.Sprintf(":%d", a.Cfg.HttpPort))

}
