package di

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (d *DI) NewAuthClientGRPC() *grpc.ClientConn {
	if d.authConn != nil {
		return d.authConn
	}

	conn, err := grpc.NewClient(
		d.Config().AuthConnectionGRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		d.Logger().Fatal(
			"Failed to create auth client",
			zap.Error(err),
		)
		return nil
	}

	d.authConn = conn

	return d.authConn
}
