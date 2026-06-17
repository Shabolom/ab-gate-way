package customParams

import (
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCustomParam(ctx echo.Context) error {
	request := new(abDto.CreateCustomParamRequest)

	err := ctx.Bind(request)
	if err != nil {
		return render.BadRequest(ctx, err)
	}

	response, err := h.abService.CreateCustomParam(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, response)
}
