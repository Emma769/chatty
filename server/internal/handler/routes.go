package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/emma769/chatty/internal/utils"
)

func (h *Handler) Register(router *chi.Mux) {
	api := chi.NewRouter()
	api.Route("/users", h.userRoutes)

	router.Mount("/api", api)
}

func (h *Handler) userRoutes(r chi.Router) {
	r.Post("/", utils.Wrap[*HandlerError](h.CreateUser))
}
