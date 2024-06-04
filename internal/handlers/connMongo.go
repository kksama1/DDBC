package handlers

import (
	"fmt"
	"github.com/kksama1/DDBC/internal/db/mongo"
	"log"
	"net/http"
)

func MongoConn(w http.ResponseWriter, r *http.Request) {
	msg := mongo.CreateConnection()
	log.Println("mconn")
	fmt.Fprintf(w, msg)
}
