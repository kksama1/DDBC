package main

import (
	"github.com/kksama1/DDBC/internal/handlers"
	"log"
	"net/http"
)

func main() {

	log.Println("server started")

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/", handlers.MongoConn)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
