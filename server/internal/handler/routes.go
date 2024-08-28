package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) ApiRoutes() http.Handler {
	e := echo.New()

	e.Use(
		middleware.RecoverWithConfig(middleware.DefaultRecoverConfig),
		middleware.LoggerWithConfig(middleware.DefaultLoggerConfig),
	)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.POST("/api/users", h.CreateUser)

	e.POST("/api/tokens/login", h.CreateToken)

	return e
}
