package grpc

import (
	"github.com/google/wire"
	common_service "github.com/tuancamtbtx/learning-go/genproto/go/common/services"
	product_service "github.com/tuancamtbtx/learning-go/genproto/go/services"
)

var GraphSet = wire.NewSet(
	NewProductServiceHandler,
	NewAuthHandler,
	wire.Bind(new(product_service.ProductServiceServer), new(*ProductServiceHandler)),
	wire.Bind(new(common_service.AuthServiceServer), new(*AuthHandler)),
)
