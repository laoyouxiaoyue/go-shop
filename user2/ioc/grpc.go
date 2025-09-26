package ioc

import (
	"google.golang.org/grpc"
	"shop/pkg/grpcx"
	"shop/user2/config"
	"shop/user2/server"
)

func InitGRPCxServer(userServer *server.UserServer) *grpcx.Server {
	newServer := grpc.NewServer()
	userServer.Register(newServer)
	return &grpcx.Server{
		Server:     newServer,
		Port:       config.Cf.App.GRPCPort,
		Name:       config.Cf.Consul.Name,
		RemoteAddr: config.Cf.Consul.RemoteAddr,
		RemotePort: config.Cf.Consul.RemotePort,
	}
}
