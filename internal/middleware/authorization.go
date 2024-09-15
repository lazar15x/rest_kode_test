package middleware

import (
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/tools"
	log "github.com/sirupsen/logrus"
)



func Authorization(s tools.DatabaseInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token = r.Header.Get("Authorization")

			if token == "" {
				log.Error(api.ErrUnauthorized)
				api.RequestErrorHandler(w, api.ErrUnauthorized)
				return
			}

			var username string = s.GetUserLoginDetails(token)
			if username == "" {
				log.Error(api.ErrUnauthorized)
				api.RequestErrorHandler(w, api.ErrUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
