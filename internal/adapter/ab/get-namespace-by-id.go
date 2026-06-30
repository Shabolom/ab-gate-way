package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) GetNamespaceByID(ctx context.Context, id int64) (*abDto.GetNamespaceByIDResponse, error) {
	response, err := a.client.GetNamespaceByID(ctx, &authv1.GetNamespaceByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	err = getNamespaceErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseNamespace := response.GetNamespace()

	return &abDto.GetNamespaceByIDResponse{
		Message: response.Message,
		Namespace: &abDto.GetNamespace{
			ID:          responseNamespace.GetId(),
			Name:        responseNamespace.GetName(),
			Description: responseNamespace.GetDescription(),
		},
	}, nil
}
