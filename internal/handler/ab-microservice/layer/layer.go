package layer

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateLayer(ctx context.Context, request *abDto.Layer) (*dto.CommonResponse, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
