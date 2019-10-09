package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listed on port 8080 %v", err)
	}
}
