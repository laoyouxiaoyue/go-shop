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
	return grpcx.NewServer(newServer, config.Cf.Server.Host, config.Cf.Server.Port, config.Cf.Consul.Host, config.Cf.Consul.Port, config.Cf.Server.Name, config.Cf.Server.Tags, config.Cf.Server.Id)

}
