package utils

import (
	"gate-way/internal/dto"
	"strings"

	"github.com/labstack/echo/v4"
)

func TokensFromHeaders(ctx echo.Context) *dto.Tokens {
	accessToken := ctx.Request().Header.Get("access-token")
	refreshToken := ctx.Request().Header.Get("refresh-token")

	return &dto.Tokens{
		AccessToken:  strings.TrimPrefix(accessToken, "Bearer "),
		RefreshToken: strings.TrimPrefix(refreshToken, "Bearer "),
	}
}
