package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) GetLayerByID(ctx context.Context, id int64) (*abDto.GetLayerByIDResponse, error) {
	response, err := a.client.GetLayerByID(ctx, &authv1.GetLayerByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	err = getLayerErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseLayer := response.GetLayer()

	return &abDto.GetLayerByIDResponse{
		Message: response.Message,
		Layer: &abDto.GetLayer{
			ID:          responseLayer.GetId(),
			NamespaceID: responseLayer.GetNamespaceId(),
			Name:        responseLayer.GetName(),
			Description: responseLayer.GetDescription(),
		},
	}, nil
}
