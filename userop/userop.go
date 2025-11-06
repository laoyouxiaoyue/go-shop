package main

import (
	"flag"
	"fmt"
	"shop/userop/internal/config"
	"shop/userop/internal/ioc"
	"shop/userop/internal/server"
	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userop.yaml", "the config file")

func main() {
	flag.Parse()

	// 初始化日志
	ioc.InitZap()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化Nacos配置（可选，如果需要从Nacos加载配置）
	if c.NacosConfig.Host != "" {
		ioc.InitNacos(&c)
	}

	// 初始化数据库
	db := ioc.InitDB(c)

	// 初始化Redis
	redis := ioc.InitRedis(c)

	// 创建服务上下文
	ctx := svc.NewServiceContext(c, db, redis)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userop.RegisterUseropServer(grpcServer, server.NewUseropServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 注册到Consul
	if c.ConsulConfig.Host != "" {
		ioc.ServerRegister(&c)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
