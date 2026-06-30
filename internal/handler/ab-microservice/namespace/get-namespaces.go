package namespace

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetNamespaces(ctx echo.Context) error {
	response, err := h.abService.GetNamespaces(ctx.Request().Context())
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
