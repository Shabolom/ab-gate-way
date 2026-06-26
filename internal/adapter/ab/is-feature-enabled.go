package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
)

func (a *Adapter) IsFeatureEnabled(ctx context.Context, id int64) (*dto.CommonResponse, error) {
	response, err := a.client.IsFeatureEnabled(ctx, &authv1.IsFeatureEnabledRequest{FeatureToggleId: id})
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
