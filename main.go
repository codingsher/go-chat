package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codingsher/go-chat/pkg/handlers"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Home())
	router.HandleFunc("GET /ws", handlers.WsEndpoint())

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	log.Println("Starting Channel listener")
	go handlers.ListenToWs()

	slog.Info("starting server", slog.String("address", "localhost:8080"), slog.String("env", "dev"), slog.String("version", "1.0.0"))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("error starting server", slog.String("error", err.Error()))
		}
	}()
	<-done

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("error shuttind down server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successful")
}
