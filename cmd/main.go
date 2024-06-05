package main

import (
	"context"
	"github.com/kksama1/DDBC/internal/db/mongodb"
	"github.com/kksama1/DDBC/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	mongodb.Client.CreateConnection(ctx)
	defer func() {
		mongodb.Client.Disconnect(ctx)
		time.Sleep(3 * time.Second)
	}()

	mongodb.Client.PingMongo(ctx)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/", handlers.Hi)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
