package di

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (d *DI) NewAbClientGRPC() *grpc.ClientConn {
	if d.abConn != nil {
		return d.abConn
	}

	conn, err := grpc.NewClient(
		d.Config().GrpcAddressAB(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		d.Logger().Fatal(
			"Failed to create ab client",
			zap.Error(err),
		)
		return nil
	}

	d.abConn = conn

	return d.abConn
}
