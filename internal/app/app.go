package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/rpcgen/v2/golang"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"net/http"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	myrpc "news-service/internal/rpc"
	"time"
)

type App struct {
	Cfg *Config
	DB  *pg.DB
	rpc zenrpc.Server
	srv *echo.Echo
}

func New(cfg *Config, dbconn *pg.DB) *App {
	repository := db.NewNewsRepo(dbconn)
	services := newsportal.NewManager(repository)

	r := myrpc.Init(services, dbconn)

	return &App{
		Cfg: cfg,
		DB:  dbconn,
		rpc: r,
	}
}

type Config struct {
	HttpPort int
	Env      string

	DB *pg.Options
}

func (a *App) Run(ctx context.Context) error {

	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := a.srv.Shutdown(stopCtx); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("--- Доступные эндпоинты ---")
	for _, route := range a.srv.Routes() {
		fmt.Printf("Метод: %-7s Путь: %s\n", route.Method, route.Path)
	}
	fmt.Println("---------------------------")

	return a.srv.Start(fmt.Sprintf(":%d", a.Cfg.HttpPort))

}

func (a *App) InitServer() {
	srv := echo.New()

	srv.Any("/rpc", echo.WrapHandler(a.rpc))
	srv.Any("/rpc/doc", echo.WrapHandler(http.HandlerFunc(zenrpc.SMDBoxHandler)))
	srv.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	gen := rpcgen.FromSMD(a.rpc.SMD())
	srv.GET("/rpc/client.go", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.GoClient(golang.Settings{})))))
	srv.GET("/rpc/client.ts", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.TSClient(nil)))))
	srv.GET("/rpc/open", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.OpenRPC("", "")))))

	a.srv = srv
}
