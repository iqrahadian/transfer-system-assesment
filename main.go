package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iqrahadian/paperid-assesment/controller"
)

func main() {

	var (
		port string = "3000"
	)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	router.Post("/disburse", controller.Disburse)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed running service without TLS. Listening on port:%s bind: address already in use", port), err)
	}

}
