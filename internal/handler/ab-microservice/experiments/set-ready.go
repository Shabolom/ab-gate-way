package experiments

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetReadyExperiment(ctx echo.Context, experimentId int64) error {
	response, err := h.abService.SetReadyExperiment(ctx.Request().Context(), experimentId)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
