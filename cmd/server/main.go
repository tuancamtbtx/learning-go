package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	service "github.com/tuancamtbtx/learning-go/genproto/go/common/services"
	product_service "github.com/tuancamtbtx/learning-go/genproto/go/services"

	grpc_handler "github.com/tuancamtbtx/learning-go/handlers/grpc"
	"google.golang.org/grpc"
)

func main() {
	var port = flag.Int("port", 8080, "The server port")
	flag.Parse()
	fmt.Println("grpc run with port", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	service.RegisterAuthServiceServer(srv, grpc_handler.NewAuthHandler())
	product_service.RegisterProductServiceServer(srv, grpc_handler.NewProductServiceHandler())
	srv.Serve(lis)
	fmt.Println("start rpc host: ", 8080)
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// eg, ctx := errgroup.WithContext(ctx)

	// defer cancel()
	// server, err := buildServer(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// eg.Go(func() error {
	// 	return server.Start(ctx)
	// })
	// defer func() {
	// 	server.Close()
	// }()
	// if err := eg.Wait(); err != nil {
	// 	fmt.Println("Run application")
	// 	os.Exit(1)
	// }

}
