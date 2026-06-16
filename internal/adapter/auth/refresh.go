package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	"gate-way/pkg/shortcut"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) Refresh(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	md := metadata.New(map[string]string{
		"authorization": tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	refreshReply, err := a.client.Refresh(ctx, &emptypb.Empty{})
	if err != nil {
		return &dto.CommonResponse{}, err
	}

	switch refreshReply.ErrInfoReason {
	case authv1.RefreshReply_STATUS_OK:
	case authv1.RefreshReply_UNSPECIFIED:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	case authv1.RefreshReply_VALIDATION_ERROR:
		return &dto.CommonResponse{}, shortcut.ErrValidation
	case authv1.RefreshReply_INVALID_REQUEST:
		return &dto.CommonResponse{}, shortcut.ErrInvalidRequest
	default:
		return &dto.CommonResponse{}, shortcut.ErrUnspecifiedRequest
	}

	return &dto.CommonResponse{
		Message: refreshReply.Message,
	}, nil
}
