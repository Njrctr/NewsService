package web

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	g "news-service/internal/global"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context, handlers *gin.Engine) {

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", g.Cfg.App.HttpPort),
		Handler:     handlers,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}
	go func() {
		<-ctx.Done()

		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(stopCtx); err != nil {

		}
	}()
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		g.Errors <- err
	}
}
