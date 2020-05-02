package grpc

import (
	"net"

	"github.com/sergiosegrera/store/cart/service"
	"github.com/sergiosegrera/store/cart/transport/grpc/bindings"
	"github.com/sergiosegrera/store/cart/transport/grpc/pb"
	"google.golang.org/grpc"
)

func Serve(svc *service.Service) error {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterCartServiceServer(server, bindings.GRPCBinding{svc})

	return server.Serve(listener)
}
