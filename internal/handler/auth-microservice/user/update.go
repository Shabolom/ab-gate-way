package userHandler

import (
	userDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/internal/render"
	"gate-way/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Update(ctx echo.Context) error {
	tokens := utils.TokensFromHeaders(ctx)

	request := new(userDto.UpdateUser)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, err)
	}

	response, err := h.authService.UpdateUser(ctx.Request().Context(), request, tokens)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
