package layer

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateLayer(ctx context.Context, request *abDto.Layer) (*dto.CommonResponse, error)
	GetLayerByID(ctx context.Context, id int64) (*abDto.GetLayerByIDResponse, error)
	GetLayers(ctx context.Context) (*abDto.GetLayersResponse, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
