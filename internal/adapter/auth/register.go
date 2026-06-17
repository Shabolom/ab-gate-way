package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (a *Adapter) Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, *dto.Tokens, error) {
	var header metadata.MD

	registerReq := &authv1.RegisterRequest{
		Mail:     register.Mail,
		Password: register.Password,
		Name:     register.Name,
		Age:      int32(register.Age),
	}

	registerReply, err := a.client.Register(ctx, registerReq, grpc.Header(&header))
	if err != nil {
		return &dto.CommonResponse{}, nil, err
	}

	accessToken := header.Get("authorization")
	refreshToken := header.Get("refresh-token")

	tokens := &dto.Tokens{}

	if len(accessToken) > 0 {
		tokens.AccessToken = accessToken[0]
	}

	if len(refreshToken) > 0 {
		tokens.RefreshToken = refreshToken[0]
	}

	switch registerReply.ErrInfoReason {
	case authv1.RegisterReply_STATUS_OK:
	case authv1.RegisterReply_UNSPECIFIED:
		return &dto.CommonResponse{}, nil, shortcut.ErrUnspecifiedRequest
	case authv1.RegisterReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, nil, shortcut.ErrValidation
	case authv1.RegisterReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, nil, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, nil, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: registerReply.Message,
	}, tokens, nil
}
