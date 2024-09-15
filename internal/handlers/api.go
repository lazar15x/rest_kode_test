package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/lazar15x/rest_kode_test/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/auth", func(router chi.Router) {
		router.Post("/login", Authenticate)
	})

	r.Route("/lk", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/notes", GetNotes)
		router.Post("/notes", CreateNotes)
	})
}