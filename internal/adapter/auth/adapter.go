package authAdapter

import (
	authv1 "gate-way/gen/proto"

	"google.golang.org/grpc"
)

type Adapter struct {
	client authv1.AccountServiceClient
}

func New(conn *grpc.ClientConn) *Adapter {
	return &Adapter{
		client: authv1.NewAccountServiceClient(conn),
	}
}
