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

	"github.com/abhinavansh18/students_api/internal/config"
	"github.com/abhinavansh18/students_api/internal/http/handlers/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//load router
	router := http.NewServeMux()
	router.HandleFunc("/api/students", student.New())

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("server started ", slog.String("address", cfg.HTTPServer.Addr))
	//fmt.Printf("Server Started: %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("FAiled to start server")
		}
	}()
	<-done
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("Error", err.Error()))
	}
	slog.Info("Server shutdown successful")
}
