package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/config"
	"shop/goods_gozero/internal/ioc"
	ioc2 "shop/goods_gozero/internal/ioc"
	"shop/goods_gozero/internal/server"
	"shop/goods_gozero/internal/svc"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {

	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	db := ioc.InitDB(c)
	ctx := svc.NewServiceContext(c, db)

	ioc.InitNacos(&c)
	fmt.Printf("%+v\n", c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, server.NewGoodsServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	ioc2.ServerRegister(&c)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}
