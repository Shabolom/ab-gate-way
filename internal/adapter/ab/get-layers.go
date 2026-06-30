package abAdapter

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetLayers(ctx context.Context) (*abDto.GetLayersResponse, error) {
	response, err := a.client.GetLayers(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	err = getLayersErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseLayers := response.GetLayers()

	var layersBody []*abDto.GetLayer
	for _, layer := range responseLayers {
		layersBody = append(layersBody, &abDto.GetLayer{
			ID:          layer.GetId(),
			NamespaceID: layer.GetNamespaceId(),
			Name:        layer.GetName(),
			Description: layer.GetDescription(),
		})
	}

	return &abDto.GetLayersResponse{
		Message: response.Message,
		Layers:  layersBody,
	}, nil
}
