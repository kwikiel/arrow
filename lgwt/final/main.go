package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var s = make([]string, 3)

fmt.Println(s)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Get("/api/fetcher", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s))
	})

	http.ListenAndServe(":3333", r)
}
