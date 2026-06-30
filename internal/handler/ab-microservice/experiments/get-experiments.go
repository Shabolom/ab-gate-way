package experiments

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetExperiments(ctx echo.Context) error {
	response, err := h.abService.GetExperiments(ctx.Request().Context())
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
