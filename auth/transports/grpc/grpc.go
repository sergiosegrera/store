package grpc

import (
	"fmt"
	"net"

	"github.com/sergiosegrera/store/auth/config"
	"github.com/sergiosegrera/store/auth/pb"
	"github.com/sergiosegrera/store/auth/service"
	"github.com/sergiosegrera/store/auth/transports/grpc/bindings"
	"google.golang.org/grpc"
)

func Serve(svc service.AuthService, conf *config.Config) error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%v", conf.GrpcPort),
	)

	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, bindings.GrpcBinding{svc})

	return server.Serve(listener)
}
