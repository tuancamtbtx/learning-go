package grpc

import (
	"context"

	product_service "github.com/tuancamtbtx/learning-go/genproto/go/services"
)

func NewProductServiceHandler() *ProductServiceHandler {
	return &ProductServiceHandler{}
}

type ProductServiceHandler struct {
}

func (p *ProductServiceHandler) Add(ctx context.Context, req *product_service.ProductServiceAddRequest) (*product_service.ProductServiceAddResponse, error) {
	return &product_service.ProductServiceAddResponse{
		Status:  true,
		Message: "created success",
	}, nil
}
