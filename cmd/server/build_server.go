// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/tuancamtbtx/learning-go/app/server"
)

func buildServer(ctx context.Context) (*server.Server, error) {
	wire.Build(server.GraphSet)
	return nil, nil
}
