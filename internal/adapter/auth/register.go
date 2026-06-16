package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"
)

func (a *Adapter) Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, error) {
	registerReq := &authv1.RegisterRequest{
		Mail:     register.Mail,
		Password: register.Password,
		Name:     register.Name,
		Age:      int32(register.Age),
	}

	registerReply, err := a.client.Register(ctx, registerReq)
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	switch registerReply.ErrInfoReason {
	case authv1.RegisterReply_STATUS_OK:
	case authv1.RegisterReply_UNSPECIFIED:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	case authv1.RegisterReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, shortcut.ErrValidation
	case authv1.RegisterReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: registerReply.Message,
	}, nil
}
