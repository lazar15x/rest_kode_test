package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lazar15x/rest_kode_test/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	services := handlers.NewHandler()
	handlers.Handler(r, services)

	fmt.Println("Starting GO API service")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Error(err)
	}
}
