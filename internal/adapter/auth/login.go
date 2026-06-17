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

func (a *Adapter) Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, *dto.Tokens, error) {
	var header metadata.MD

	loginReq := &authv1.LoginRequest{
		Mail:     login.Mail,
		Password: login.Password,
	}

	loginReply, err := a.client.Login(
		ctx,
		loginReq,
		grpc.Header(&header),
	)
	if err != nil {
		return nil, nil, err
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

	switch loginReply.ErrInfoReason {
	case authv1.LoginReply_STATUS_OK:

	case authv1.LoginReply_UNSPECIFIED:
		return nil, nil, shortcut.ErrUnspecifiedRequest

	case authv1.LoginReply_VALIDATION_ERROR:
		return nil, nil, shortcut.ErrValidation

	case authv1.LoginReply_INVALID_REQUEST:
		return nil, nil, shortcut.ErrInvalidRequest

	default:
		return nil, nil, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: loginReply.Message,
	}, tokens, nil
}
