package server

import (
	"fmt"
	"net"

	product_service "github.com/tuancamtbtx/learning-go/genproto/go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(
	productService product_service.ProductServiceServer,
) (*GrpcServer, error) {
	srv := grpc.NewServer()

	product_service.RegisterProductServiceServer(srv, productService)
	reflection.Register(srv)

	return &GrpcServer{
		server: srv,
		addr:   fmt.Sprintf(":%v", 8000),
	}, nil
}

type GrpcServer struct {
	server *grpc.Server
	addr   string
}

func (s *GrpcServer) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *GrpcServer) Close() error {
	s.server.GracefulStop()
	return nil
}
