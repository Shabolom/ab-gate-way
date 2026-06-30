package abAdapter

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetCustomParams(ctx context.Context) (*abDto.GetCustomParamsResponse, error) {
	response, err := a.client.GetCustomParams(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	err = getCustomParamsErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseCustomParams := response.GetCustomParams()

	var customParamsBody []*abDto.GetCustomParam
	for _, customParam := range responseCustomParams {
		customParamsBody = append(customParamsBody, &abDto.GetCustomParam{
			ID:          customParam.GetId(),
			NamespaceID: customParam.GetNamespaceId(),
			Name:        customParam.GetName(),
			Type:        customParam.GetType(),
		})
	}

	return &abDto.GetCustomParamsResponse{
		Message:      response.Message,
		CustomParams: customParamsBody,
	}, nil
}
