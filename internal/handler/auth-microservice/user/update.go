package user

import (
	userDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateUsers(ctx echo.Context) error {
	request := new(userDto.UpdateUser)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, err)
	}

	response, err := h.authService.UpdateUser(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
