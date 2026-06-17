package namespace

import (
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateNamespace(ctx echo.Context) error {
	request := new(abDto.Namespace)

	err := ctx.Bind(request)
	if err != nil {
		return render.BadRequest(ctx, err)
	}

	response, err := h.abService.CreateNamespace(ctx.Request().Context(), request)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return render.JSON(ctx, http.StatusCreated, response)
}
