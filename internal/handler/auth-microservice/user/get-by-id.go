package userHandler

import (
	"gate-way/internal/render"
	"gate-way/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetByID(ctx echo.Context) error {
	tokens := utils.TokensFromHeaders(ctx)

	response, err := h.authService.GetUser(ctx.Request().Context(), tokens)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
