package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	"gate-way/pkg/shortcut"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) DeleteUsers(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	md := metadata.New(map[string]string{
		"authorization": tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	deleteReply, err := a.client.DeleteUsers(ctx, &emptypb.Empty{})
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	switch deleteReply.ErrInfoReason {
	case authv1.DeleteUsersReply_STATUS_OK:
	case authv1.DeleteUsersReply_UNSPECIFIED:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	case authv1.DeleteUsersReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, shortcut.ErrValidation
	case authv1.DeleteUsersReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: deleteReply.Message,
	}, nil
}
