package feature

import (
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateFeatureToggle(ctx echo.Context) error {
	request := new(abDto.Feature)

	err := ctx.Bind(request)
	if err != nil {
		return render.BadRequest(ctx, err.(error))
	}

	response, err := h.abService.CreateFeature(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusCreated, response)
}
