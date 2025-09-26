package user2

import (
	"fmt"
	"os"
	"os/signal"
	"shop/user2/config"
	"shop/user2/ioc"
	"shop/user2/repository"
	"shop/user2/repository/dao"
	"shop/user2/server"
	"shop/user2/service"
	"syscall"
)

func main() {
	config.Cf, _ = config.Load()
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
	err := s.Serve()
	if err != nil {
		fmt.Println(err)
		panic("error")
	}

}
