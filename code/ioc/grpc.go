package ioc

import (
	"google.golang.org/grpc"
	"shop/code/config"
	"shop/code/server"
	"shop/pkg/grpcx"
)

func InitGRPCxServer(codeServer *server.CodeServer) *grpcx.Server {
	newServer := grpc.NewServer()
	codeServer.Register(newServer)

	return &grpcx.Server{
		Server:     newServer,
		Port:       config.Cf.App.GRPCPort,
		Name:       config.Cf.Consul.Name,
		RemoteAddr: config.Cf.Consul.RemoteAddr,
		RemotePort: config.Cf.Consul.RemotePort,
	}
}
