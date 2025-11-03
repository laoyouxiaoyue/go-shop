package main

import (
	"flag"
	"fmt"
	"shop/inventory/internal/ioc"

	"shop/inventory/internal/config"
	"shop/inventory/internal/server"
	"shop/inventory/internal/svc"
	"shop/inventory/inventory"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/inventory.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	db := ioc.InitDB(c)
	redis := ioc.InitRedis(c)
    rs := ioc.InitRedsync(redis)
    ctx := svc.NewServiceContext(c, db, redis, rs)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		inventory.RegisterInventoryServer(grpcServer, server.NewInventoryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
