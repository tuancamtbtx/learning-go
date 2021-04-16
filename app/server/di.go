package server

import (
	"github.com/google/wire"
	grpc_handler "github.com/tuancamtbtx/learning-go/handlers/grpc"
)

var Deps = wire.NewSet(
	grpc_handler.GraphSet,
)

var GraphSet = wire.NewSet(
	Deps,
	NewGrpcServer,
	NewServer,
)
