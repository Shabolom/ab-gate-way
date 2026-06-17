package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) UserExperiment(ctx context.Context, request *abDto.UserExperimentRequest) (*abDto.ExperimentsReply, error) {
	experimentsResp := new(abDto.ExperimentsReply)

	requestBody := &authv1.ExperimentRequest{
		SplitId:   request.SplitID,
		DeviceId:  request.DeviceID,
		Namespace: request.Namespace,
		City:      request.City,
		Store:     request.Store,
	}
	for _, param := range request.Params {
		requestBody.Params = append(requestBody.Params, &authv1.Param{
			ParamName: param.ParamName,
			Value:     param.Value,
		})
	}

	response, err := a.client.UserExperiment(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = experimentsReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	for _, experimentReply := range response.ExperimentsReply {
		experimentsResp.ExperimentsReply = append(experimentsResp.ExperimentsReply, abDto.ExperimentReply{
			ExperimentName: experimentReply.ExperimentName,
			GroupName:      experimentReply.GroupName,
		})
	}

	return experimentsResp, nil
}
