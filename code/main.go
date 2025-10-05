package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"shop/code/config"
	"shop/code/ioc"
	"shop/code/repository"
	"shop/code/repository/cache"
	"shop/code/server"
	"shop/code/service"
	"syscall"
	"time"
)

func main() {
	ioc.InitZap()
	var err error
	config.Cf, err = config.Load()
	ioc.InitNacos()
	if err != nil {
		zap.L().Panic("配置文件获取失败", zap.Error(err))
	}
	tm := ioc.InitYamlTemplateManager()
	cmd := ioc.InitRedis()
	ruc := cache.NewRedisCodeCache(cmd, time.Minute)
	repo := repository.NewCodeRepository(ruc)
	svc := service.NewCodeService(repo, tm)
	svr := server.NewCodeServer(svc)
	s := ioc.InitGRPCxServer(svr)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		s.DeRegister()
	}()
	err = s.Serve()
	if err != nil {
		fmt.Println(err)
		panic("error")
	}

}
