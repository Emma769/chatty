package middleware

import (
	"context"
	"net/http"
	"time"
)

type infolog interface {
	InfoContext(context.Context, string, ...any)
}

func Logger(lg infolog) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			lg.InfoContext(
				context.Background(),
				"incoming request",
				"method",
				r.Method,
				"path",
				r.RequestURI,
				"latency",
				time.Since(start),
			)
		})
	}
}
