package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) GetCustomParamByID(ctx context.Context, id int64) (*abDto.GetCustomParamByIDResponse, error) {
	response, err := a.client.GetCustomParamByID(ctx, &authv1.GetCustomParamByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	err = getCustomParamErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseCustomParam := response.GetCustomParam()

	return &abDto.GetCustomParamByIDResponse{
		Message: response.Message,
		CustomParam: &abDto.GetCustomParam{
			ID:          responseCustomParam.GetId(),
			NamespaceID: responseCustomParam.GetNamespaceId(),
			Name:        responseCustomParam.GetName(),
			Type:        responseCustomParam.GetType(),
		},
	}, nil
}
