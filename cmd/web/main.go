package main

import (
	"html/template"
	"log"
	"net/http"

	"dfcsr/internal/dog"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	tpl := template.Must(template.ParseGlob("templates/*"))

	dogRepo := dog.NewMemoryRepository()
	dogService := dog.NewService(dogRepo)
	dogController := dog.NewController(dogService, tpl)

	mux.Get("/dog", dogController.ByName)
	mux.Get("/dog/all", dogController.All)

	log.Println("listening on port 8080")
	_ = http.ListenAndServe(":8080", mux)
}
