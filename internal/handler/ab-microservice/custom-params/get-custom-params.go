package customParams

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCustomParams(ctx echo.Context) error {
	response, err := h.abService.GetCustomParams(ctx.Request().Context())
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
