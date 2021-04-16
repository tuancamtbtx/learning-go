package server

import (
	"context"

	"github.com/hienduyph/goss/utils/shutdowns"
	"golang.org/x/sync/errgroup"
)

func NewServer(
	grpcServer *GrpcServer,
) *Server {
	return &Server{
		grpcServer: grpcServer,
	}
}

type Server struct {
	grpcServer *GrpcServer
}

func (s *Server) Start(parentCtx context.Context) error {
	eg, ctx := errgroup.WithContext(parentCtx)
	eg.Go(func() error {

		return shutdowns.BlockListen(ctx, s.grpcServer)
	})

	return eg.Wait()
}

func (s *Server) Close() error {
	// the `BlockListen` has already handled graceful shutdown when ctx is cancled.
	// just do nothing
	return nil
}
