package handlers

import (
	"fmt"
	"net/http"
)

func Hi(w http.ResponseWriter, r *http.Request) {
	//mongodb.Client.PingMongo(context.TODO())
	fmt.Fprintf(w, "hi")
}
