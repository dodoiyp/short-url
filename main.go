package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"short-url/initialize"
	"short-url/pkg/global"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("application start failed %v :%v", err, string(debug.Stack()))
		}
	}()

	initialize.Config()
	initialize.Mysql()
	initialize.Cache()
	r := initialize.Routers()

	host := "0.0.0.0"
	port := global.Conf.System.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen error: %v", err)
		}
	}()

	log.Printf("server is running at %s:%d", host, port)

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("server forced to shutdown: %v ", err)
	}

	log.Printf("server exiting")
}
