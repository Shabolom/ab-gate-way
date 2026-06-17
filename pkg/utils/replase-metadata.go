package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func ReplaceTokensMetadata(ctx context.Context, accessToken, refreshToken string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	md.Set("authorization", accessToken)
	md.Set("refresh-token", refreshToken)

	return metadata.NewOutgoingContext(ctx, md)
}
