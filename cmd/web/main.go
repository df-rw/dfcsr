package main

import (
	"fmt"
	"net/http"

	"dfcsr/internal/dog"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	dogRepo := dog.NewMemoryRepository()
	dogService := dog.NewService(dogRepo)
	dogController := dog.NewController(dogService)

	mux.Get("/dog/all", dogController.All)

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
