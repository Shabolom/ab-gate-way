package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetUser(ctx context.Context, tokens *dto.Tokens) (*authMicroserviceDto.User, error) {
	md := metadata.New(map[string]string{
		"authorization": tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	userReply, err := a.client.GetUser(ctx, &emptypb.Empty{})
	if err != nil {
		return &authMicroserviceDto.User{}, err
	}

	switch userReply.ErrInfoReason {
	case authv1.GetUserReply_STATUS_OK:
	case authv1.GetUserReply_UNSPECIFIED:
		return &authMicroserviceDto.User{}, shortcut.ErrUnspecifiedRequest
	case authv1.GetUserReply_VALIDATION_ERROR:
		return &authMicroserviceDto.User{}, shortcut.ErrValidation
	case authv1.GetUserReply_INVALID_REQUEST:
		return &authMicroserviceDto.User{}, shortcut.ErrInvalidRequest
	default:
		return &authMicroserviceDto.User{}, shortcut.ErrUnspecifiedRequest
	}

	user := userReply.GetUser()
	if user == nil {
		return &authMicroserviceDto.User{}, shortcut.ErrUnspecifiedResponseGetUser
	}

	return &authMicroserviceDto.User{
		ID:        user.GetId(),
		Mail:      user.GetMail(),
		Name:      user.GetName(),
		Age:       int(user.GetAge()),
		CreatedAT: user.GetCreatedAt().AsTime(),
		UpdatedAT: user.GetAddedAt().AsTime(),
	}, nil
}
