package middleware

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/emma769/chatty/internal/utils"
)

type errorlog interface {
	ErrorContext(context.Context, string, ...any)
}

func Recover(lg errorlog) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					_ = utils.WriteJsonE(
						w,
						http.StatusInternalServerError,
						map[string]string{"detail": "internal server error"},
					)

					lg.ErrorContext(
						context.Background(),
						"server error",
						"detail",
						fmt.Sprintf("%v", err),
						"trace",
						string(debug.Stack()),
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
