package main

import (
	"flag"
	"fmt"
	"shop/order/internal/ioc"

	"shop/order/internal/config"
	"shop/order/internal/server"
	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	db := ioc.InitDB(c)
	redis := ioc.InitRedis(c)
	goodsCli, err := ioc.InitGoodsClient(c)
	if err != nil {
		panic(err)
	}
	invCli, err := ioc.InitInventoryClient(c)
	if err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c, db, redis, goodsCli, invCli)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	ioc.ServerRegister(&c)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
