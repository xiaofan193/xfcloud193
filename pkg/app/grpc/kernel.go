package grpc

import (
	"github.com/xiaofan193/xifancloud193/internal/framework"
	pkgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcEngine(container framework.Container) (*pkgrpc.Server, error) {
	s := pkgrpc.NewServer()
	// here register service
	// todo protoc gen code

	reflection.Register(s)
	return s, nil
}
