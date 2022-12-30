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
	port := os.Getenv("PORT")
	if port == "" {
		port = "9726"
	}
	fmt.Println("start server ... ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
