package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"shop/user/config"
	"shop/user/ioc"
	"shop/user/repository"
	"shop/user/repository/dao"
	"shop/user/server"
	"shop/user/service"
	"syscall"
)

func main() {
	ioc.InitZap()
	var err error
	config.Cf, err = config.Load()
	if err != nil {
		zap.S().Error("获取配置文件失败", zap.Error(err))
		return
	}

	db := ioc.InitDB()
	userDao := dao.NewGormUserDao(db)
	repo := repository.NewGormUserRepository(userDao)
	svc := service.NewUserService(repo)
	svr := server.NewUserServer(svc)
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
