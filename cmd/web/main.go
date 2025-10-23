package main

import (
	"fmt"
	"net/http"

	"dfcsr/internal/dog"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	dogController := dog.NewController(dog.NewService(dog.NewMemoryRepository()))

	mux.Get("/dog", dogController.ByName)
	mux.Get("/dog/all", dogController.All)

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
