package feature

import (
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetFeatureToggleStatus(ctx echo.Context, featureToggleId int64) error {
	request := &abDto.SetFeatureStatus{
		FeatureID: featureToggleId,
	}

	err := ctx.Bind(&request)
	if err != nil {
		return render.BadRequest(ctx, err)
	}

	response, err := h.abService.SetFeatureToggleStatus(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
