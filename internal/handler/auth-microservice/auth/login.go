package auth

import (
	"gate-way/internal/dto/auth-microservice-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(ctx echo.Context) error {
	request := new(authMicroserviceDto.Login)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, err)
	}

	response, tokens, err := h.authService.Login(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	ctx.Response().Header().Add("authorization", tokens.AccessToken)
	ctx.Response().Header().Add("refresh-token", tokens.RefreshToken)
	return render.JSON(ctx, http.StatusOK, response)
}
