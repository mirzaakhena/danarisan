package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mirzaakhena/danarisan/infrastructure/log"
)

// GracefullyShutdown will handle http server with gracefully shutdown mechanism
type GracefullyShutdown struct {
	httpServer *http.Server
}

func NewGracefullyShutdown(handler http.Handler, address string) GracefullyShutdown {
	return GracefullyShutdown{
		httpServer: &http.Server{
			Addr:    address,
			Handler: handler,
		},
	}
}

// RunWithGracefullyShutdown is ...
func (r *GracefullyShutdown) RunWithGracefullyShutdown() {

	go func() {
		if err := r.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(context.Background(), "listen: %s", err)
			os.Exit(1)
		}
	}()

	log.Info(context.Background(), "server is running ...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(context.Background(), "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.httpServer.Shutdown(ctx); err != nil {
		log.Error(context.Background(), "Server forced to shutdown:", err)
		os.Exit(1)
	}

	log.Info(context.Background(), "Server stoped.")

}
