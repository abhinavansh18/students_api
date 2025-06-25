package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhinavansh18/students_api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//load router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	//setup server
	server = http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	err := server.ListenandServe()
	if err != nil {
		log.Fatal("FAiled to start server")
	}
	fmt.Println("Server Started")
}
