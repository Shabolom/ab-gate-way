package abAdapter

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetNamespaces(ctx context.Context) (*abDto.GetNamespacesResponse, error) {
	response, err := a.client.GetNamespaces(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	err = getNamespacesErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseNamespaces := response.GetNamespaces()

	var namespacesBody []*abDto.GetNamespace
	for _, namespace := range responseNamespaces {
		namespacesBody = append(namespacesBody, &abDto.GetNamespace{
			ID:          namespace.GetId(),
			Name:        namespace.GetName(),
			Description: namespace.GetDescription(),
		})
	}

	return &abDto.GetNamespacesResponse{
		Message:    response.Message,
		Namespaces: namespacesBody,
	}, nil
}
