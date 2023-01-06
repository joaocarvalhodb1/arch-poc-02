package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewgRPCConnection(gRPCService string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		gRPCService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
