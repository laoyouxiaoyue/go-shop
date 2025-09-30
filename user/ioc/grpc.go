package ioc

import (
	"google.golang.org/grpc"
	"shop/pkg/grpcx"
	"shop/user/config"
	"shop/user/server"
)

func InitGRPCxServer(userServer *server.UserServer) *grpcx.Server {
	newServer := grpc.NewServer()
	userServer.Register(newServer)
	return &grpcx.Server{
		Server:     newServer,
		Port:       config.Cf.Server.Port,
		Name:       config.Cf.Server.Name,
		RemoteAddr: config.Cf.Consul.Host,
		RemotePort: config.Cf.Consul.Port,
	}
}
