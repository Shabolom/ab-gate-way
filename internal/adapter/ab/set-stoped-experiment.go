package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
)

func (a *Adapter) SetStopedExperiment(ctx context.Context, request int64) (*dto.CommonResponse, error) {
	requestBody := &authv1.SetStopedExperimentRequest{ExperimentId: request}

	response, err := a.client.SetStopedExperiment(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
