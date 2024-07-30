package handler

import (
	"errors"
	"net/http"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/services"
	"github.com/emma769/chatty/internal/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var in model.UserIn

	if err := utils.ReadJson(w, r, &in); err != nil {
		return utils.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := in.Validate(); err != nil {
		return utils.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	user, err := h.User.Create(r.Context(), in)

	if err != nil && errors.Is(err, services.ErrDuplicateKey) {
		return utils.NewError(http.StatusConflict, "email already in use")
	}

	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusCreated, user)
}
