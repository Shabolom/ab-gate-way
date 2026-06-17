package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) CreateLayer(ctx context.Context, layer *abDto.Layer) (*dto.CommonResponse, error) {
	requestBody := &authv1.CreateLayerRequest{
		NamespaceId: layer.NamespaceID,
		Name:        layer.Name,
		Description: layer.Description,
	}

	response, err := a.client.CreateLayer(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
