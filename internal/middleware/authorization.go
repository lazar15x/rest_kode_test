package middleware

import (
	"errors"
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/tools"
	log "github.com/sirupsen/logrus"
)

var  ErrUnauthorized = errors.New("unauthorized access")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("Authorization")
		var err error
		if token == "" {
			log.Error(ErrUnauthorized )
			api.RequestErrorHandler(w, ErrUnauthorized )
			return
		}

		var db *tools.DatabaseInterface
		if db, err = tools.NewDatabase(); err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var username string = (*db).GetUserLoginDetails(token)
		if username == "" {
			log.Error(ErrUnauthorized )
			api.RequestErrorHandler(w, ErrUnauthorized )
			return
		}

		next.ServeHTTP(w, r)
	})
}