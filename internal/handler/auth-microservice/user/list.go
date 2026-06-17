package user

import (
	"gate-way/internal/render"
	"gate-way/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersList(ctx echo.Context) error {
	tokens := utils.TokensFromHeaders(ctx)

	response, err := h.authService.GetUsersList(ctx.Request().Context(), tokens)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
