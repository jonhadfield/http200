package main

import (
	"log"
	"net/http"

	"github.com/jonhadfield/http200"
)

func main() {
	handler := http.HandlerFunc(http200.Server)
	if err := http.ListenAndServe(":80", handler); err != nil {
		log.Fatalf("could not listen on port 80 %v", err)
	}
}
