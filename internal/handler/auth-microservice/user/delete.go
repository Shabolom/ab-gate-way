package user

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteUsers(ctx echo.Context) error {
	response, err := h.authService.DeleteUser(ctx.Request().Context())
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
