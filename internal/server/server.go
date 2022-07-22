package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer http.Server
}

func NewApp(handler http.Handler) *App {
	port := os.Getenv("PORT")

	return &App{
		httpServer: http.Server{
			Addr:    port,
			Handler: handler,
		},
	}
}

func (a *App) Run() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func(s chan os.Signal) {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Println(err.Error())
			done <- os.Interrupt
		}
	}(done)

	<-done
	defer close(done)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	log.Println("Server gracefully closed")
	return a.httpServer.Shutdown(ctx)
}
