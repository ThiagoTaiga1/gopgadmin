package main

import (
	"fmt"
	"main/configs"
	"main/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServePort()), r)
}
