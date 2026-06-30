package namespace

import (
	"gate-way/internal/render"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetNamespaceByID(ctx echo.Context, id int64) error {
	response, err := h.abService.GetNamespaceByID(ctx.Request().Context(), id)
	if err != nil {
		return render.FromError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
