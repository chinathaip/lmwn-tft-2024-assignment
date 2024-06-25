package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
}

func NewGinRouter() *GinRouter {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())

	return &GinRouter{e}
}

func (r *GinRouter) Start(port string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: r}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("error while shutting down the server: %v", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start listener: %v", err)
	}
}
