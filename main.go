package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bytemate/lark-github-bot/src"
)

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		src.LarkServer.EventCallback.ListenCallback(r.Context(), r.Body, w)
	})
	fmt.Println("start server ... 9726")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9726"
	}
	log.Fatal(http.ListenAndServe(":9726", nil))
}
