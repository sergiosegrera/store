package grpc

import (
	"fmt"
	"net"

	"github.com/sergiosegrera/store/cart/config"
	"github.com/sergiosegrera/store/cart/pb"
	"github.com/sergiosegrera/store/cart/service"
	"github.com/sergiosegrera/store/cart/transports/grpc/bindings"
	"google.golang.org/grpc"
)

func Serve(svc service.CartService, conf *config.Config) error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%v", conf.GrpcPort),
	)

	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterCartServiceServer(server, bindings.GrpcBinding{svc})

	return server.Serve(listener)
}
