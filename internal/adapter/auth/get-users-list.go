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

func (a *Adapter) GetUsersList(ctx context.Context, tokens *dto.Tokens) ([]*authMicroserviceDto.User, error) {
	md := metadata.New(map[string]string{
		"authorization": tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	usersReply, err := a.client.GetUsersList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	switch usersReply.ErrInfoReason {
	case authv1.GetUsersReply_STATUS_OK:
	case authv1.GetUsersReply_UNSPECIFIED:
		return nil, shortcut.ErrUnspecifiedRequest
	case authv1.GetUsersReply_VALIDATION_ERROR:
		return nil, shortcut.ErrValidation
	case authv1.GetUsersReply_INVALID_REQUEST:
		return nil, shortcut.ErrInvalidRequest
	default:
		return nil, shortcut.ErrUnspecifiedRequest
	}

	users := make([]*authMicroserviceDto.User, 0, len(usersReply.Users))
	for _, user := range usersReply.Users {
		users = append(users, &authMicroserviceDto.User{
			ID:        user.GetId(),
			Mail:      user.GetMail(),
			Name:      user.GetName(),
			Age:       int(user.GetAge()),
			CreatedAT: user.GetCreatedAt().AsTime(),
			UpdatedAT: user.GetAddedAt().AsTime(),
		})
	}

	return users, nil
}
