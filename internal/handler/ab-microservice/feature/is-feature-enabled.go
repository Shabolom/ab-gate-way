package feature

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IsFeatureEnabled(ctx echo.Context, featureToggleId int64) error {
	response, err := h.abService.IsFeatureEnabled(ctx.Request().Context(), featureToggleId)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
