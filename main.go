package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iqrahadian/paperid-assesment/controller"
	"github.com/iqrahadian/paperid-assesment/event"
	"github.com/iqrahadian/paperid-assesment/repo"
)

func main() {

	var (
		port string = "3000"
		wg   sync.WaitGroup
	)

	wg.Add(1)
	go event.StartConsumer(&wg)

	// init base data
	repo.OnLoad()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	router.Post("/disburse", controller.Disburse)
	router.Get("/data", controller.ShowData)

	fmt.Println("Serving api at post %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed running service without TLS. Listening on port:%s bind: address already in use", port), err)
	}

	wg.Wait()

}
