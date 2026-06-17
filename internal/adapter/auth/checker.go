package authAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/pkg/shortcut"
	"gate-way/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) Checker(ctx context.Context) (context.Context, error) {
	var header metadata.MD

	response, err := a.client.Check(
		ctx,
		&emptypb.Empty{},
		grpc.Header(&header),
	)
	if err != nil {
		return ctx, err
	}

	switch response.ErrInfoReason {
	case authv1.CheckReply_STATUS_OK:
	case authv1.CheckReply_ACCESS_TOKEN_EXPIRED:
		return ctx, shortcut.ErrAccessTokenExpired
	case authv1.CheckReply_REFRESH_TOKEN_EXPIRED:
		return ctx, shortcut.ErrRefreshTokenExpired
	case authv1.CheckReply_ACCESS_TOKEN_NOT_FOUND:
		return ctx, shortcut.ErrAccessTokenNotFound
	case authv1.CheckReply_REFRESH_TOKEN_NOT_FOUND:
		return ctx, shortcut.ErrRefreshTokenNotFound
	case authv1.CheckReply_ACCESS_TOKEN_REVOKED:
		return ctx, shortcut.ErrAccessTokenRevoked
	case authv1.CheckReply_REFRESH_TOKEN_REVOKED:
		return ctx, shortcut.ErrRefreshTokenRevoked
	case authv1.CheckReply_INVALID_ACCESS_TOKEN:
		return ctx, shortcut.ErrInvalidAccessToken
	case authv1.CheckReply_INVALID_REFRESH_TOKEN:
		return ctx, shortcut.ErrInvalidRefreshToken
	case authv1.CheckReply_TOKEN_PAIR_MISMATCH:
		return ctx, shortcut.ErrTokenPairMismatch
	case authv1.CheckReply_SESSION_NOT_FOUND:
		return ctx, shortcut.ErrSessionNotFound
	case authv1.CheckReply_VALIDATION_ERROR:
		return ctx, shortcut.ErrValidation
	case authv1.CheckReply_INVALID_REQUEST:
		return ctx, shortcut.ErrInvalidRequest
	case authv1.CheckReply_UNSPECIFIED:
		return ctx, shortcut.ErrUnspecifiedRequest
	}

	accessToken := utils.GetFromMetadata(header, "authorization")
	refreshToken := utils.GetFromMetadata(header, "refresh-token")

	if accessToken == "" || refreshToken == "" {
		return ctx, nil
	}

	ctx = utils.ReplaceTokensMetadata(ctx, accessToken, refreshToken)

	return ctx, nil
}
