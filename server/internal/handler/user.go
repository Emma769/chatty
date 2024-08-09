package handler

import (
	"errors"
	"net/http"

	"github.com/emma769/chatty/internal/model"
	"github.com/emma769/chatty/internal/services"
	"github.com/emma769/chatty/internal/utils"
	"github.com/emma769/chatty/pkg/validator"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var in model.UserIn

	if err := utils.ReadJSON(w, r, &in); err != nil {
		return NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.ValidateStruct(in); err != nil {
		return NewError(http.StatusUnprocessableEntity, err.Error())
	}

	user, err := h.User.Create(r.Context(), in)

	if err != nil && errors.Is(err, services.ErrDuplicateKey) {
		return NewError(http.StatusConflict, "email already in use")
	}

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, user)
}
