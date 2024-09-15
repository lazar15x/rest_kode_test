package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/lazar15x/rest_kode_test/internal/middleware"
	"github.com/lazar15x/rest_kode_test/internal/tools"
	log "github.com/sirupsen/logrus"
)

type Services struct {
	db tools.DatabaseInterface
}

// Инициализация базы данных
func NewHandler() *Services {
	db, err := tools.NewDatabase()
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных %v", err)
	}
	return &Services{db: db}
}

func Handler(r *chi.Mux, s *Services) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/auth", func(router chi.Router) {
		router.Post("/login", s.Authentication)
	})

	r.Route("/lk", func(router chi.Router) {
		router.Use(middleware.Authorization(s.db))

		router.Get("/notes", s.GetNotes)
		router.Post("/notes", s.CreateNotes)
	})
}
