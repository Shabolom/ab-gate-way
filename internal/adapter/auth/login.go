package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"
)

func (a *Adapter) Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, error) {
	loginReq := &authv1.LoginRequest{
		Mail:     login.Mail,
		Password: login.Password,
	}

	loginReply, err := a.client.Login(ctx, loginReq)
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	switch loginReply.ErrInfoReason {
	case authv1.LoginReply_STATUS_OK:
	case authv1.LoginReply_UNSPECIFIED:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	case authv1.LoginReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, shortcut.ErrValidation
	case authv1.LoginReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: loginReply.Message,
	}, nil
}
