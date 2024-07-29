package middleware

import (
	"net/http"
	"strings"
)

type CorsOptions struct {
	AllowOrigins, AllowMethods, AllowHeaders []string
}

func EnableCORS(opts *CorsOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Vary", "Origin")
			origin := r.Header.Get("Origin")

			w.Header().Add("Vary", "Access-Control-Request-Method")
			method := r.Header.Get("Access-Control-Request-Method")

			for i := range opts.AllowOrigins {
				if origin == opts.AllowOrigins[i] {
					w.Header().Set("Access-Control-Allow-Origin", origin)

					if method != "" && r.Method == http.MethodOptions {
						methods := strings.Join(opts.AllowMethods, ", ")
						w.Header().Set("Access-Control-Allow-Methods", methods)

						headers := strings.Join(opts.AllowHeaders, ", ")
						w.Header().Set("Access-Control-Allow-Headers", headers)

						w.WriteHeader(http.StatusOK)
						return
					}

					break
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
