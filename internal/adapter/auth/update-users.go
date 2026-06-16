package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"google.golang.org/grpc/metadata"
)

func (a *Adapter) UpdateUsers(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	md := metadata.New(map[string]string{
		"authorization": tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &authv1.UpdateUser{
		Mail: updateUser.Mail,
		Name: updateUser.Name,
		Age:  uint32(updateUser.Age),
	}

	updateReq := &authv1.UpdateUsersRequest{
		UpdatedUser: data,
	}

	updateReply, err := a.client.UpdateUsers(ctx, updateReq)
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	switch updateReply.ErrInfoReason {
	case authv1.UpdateUsersReply_STATUS_OK:
	case authv1.UpdateUsersReply_UNSPECIFIED:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	case authv1.UpdateUsersReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, shortcut.ErrValidation
	case authv1.UpdateUsersReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: updateReply.Message,
	}, nil
}
