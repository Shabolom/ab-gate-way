package feature

import (
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IsUserInFeature(ctx echo.Context) error {
	request := new(abDto.UserFeatureReq)

	err := ctx.Bind(request)
	if err != nil {
		return render.BadRequest(ctx, err.(error))
	}

	response, err := h.abService.IsUserInFeature(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
