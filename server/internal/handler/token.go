package handler

import (
	"cmp"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/emma769/chatty/internal/data"
	"github.com/emma769/chatty/internal/services"
	"github.com/emma769/chatty/pkg/passlib"
	"github.com/emma769/chatty/pkg/validator"
)

func (h *Handler) CreateToken(c echo.Context) error {
	var in data.TokenIn

	if err := c.Bind(&in); err != nil {
		return err
	}

	v := validator.New()

	if errs := v.ValidateStruct(in); errs != nil {
		return c.JSON(http.StatusUnprocessableEntity, errs)
	}

	user, err := h.User.FindByEmail(h.ctx, in.Email)
	if err != nil && errors.Is(err, services.ErrNotFound) {
		return echo.ErrUnauthorized
	}
	if err != nil {
		return err
	}

	if !passlib.IsMatch(in.Password, user.Password) {
		return echo.ErrUnauthorized
	}

	accesstoken, accessexp, err := h.Token.Encrypt(
		user.UserID,
		cmp.Or(h.Getter.GetDuration("JWT_ACCESS_EXPIRE"), 15*time.Minute),
	)
	if err != nil {
		return err
	}

	refreshtoken, refreshexp, err := h.Session.Create(h.ctx, data.SessionIn{
		UserID:   user.UserID,
		Email:    user.Email,
		Scope:    data.Scope_Authentication,
		ValidFor: cmp.Or(h.Getter.GetDuration("JWT_REFRESH_EXPIRE"), 30*time.Minute),
	})
	if err != nil {
		return err
	}

	accessToken := data.AccessToken{
		Value:     accesstoken,
		ValidTill: accessexp,
	}

	refreshToken := &data.RefreshToken{
		Value:     refreshtoken,
		ValidTill: refreshexp,
	}

	return c.JSON(http.StatusOK, createToken(user, accessToken, refreshToken))
}

func createToken(
	user *data.User,
	accesstoken data.AccessToken,
	refreshtoken *data.RefreshToken,
) data.Token {
	return data.Token{
		User:         user,
		AccessToken:  accesstoken,
		RefreshToken: refreshtoken,
	}
}
