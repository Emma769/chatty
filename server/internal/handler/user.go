package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/emma769/chatty/internal/data"
	"github.com/emma769/chatty/internal/services"
	"github.com/emma769/chatty/pkg/validator"
)

func (h *Handler) CreateUser(c echo.Context) error {
	var in data.UserIn

	if err := c.Bind(&in); err != nil {
		return err
	}

	v := validator.New()

	if errs := v.ValidateStruct(in); errs != nil {
		return c.JSON(http.StatusUnprocessableEntity, errs)
	}

	user, err := h.User.Create(h.ctx, in)
	if err != nil && errors.Is(err, services.ErrDuplicateKey) {
		return echo.NewHTTPError(http.StatusBadRequest, "email already in use")
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
