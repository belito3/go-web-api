package app

import (
	"context"
	"fmt"
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/app/route"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"go.uber.org/dig"
	"net/http"
	"time"
)

func InitHTTPServer(ctx context.Context, container *dig.Container) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	srv := &http.Server{
		Addr: addr,
		Handler: route.InitGinEngine(container),
		//ReadTimeout: 5 * time.Second,
		//WriteTimeout: 10 * time.Second,
		//IdleTimeout: 15 * time.Second,
	}

	go func() {
		logger.Printf(ctx, "HTTP server is running at %s.", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		// Wait for interrupt signal to gracefully shutdown the app with
		// a timeout
		ctx, cancel := context.WithTimeout(ctx, time.Second * time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}