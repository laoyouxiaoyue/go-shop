package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"shop/user/global"
	"shop/user/handler"
	"shop/user/proto"
	service2 "shop/user/service"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip address")
	Port := flag.Int("port", 5005, "port number")

	flag.Parse()
	fmt.Println("IP:", *IP, "Port:", *Port)
	server := grpc.NewServer()
	serviceV := service2.NewUserService(global.DB)
	proto.RegisterUserServer(server, handler.NewUserServer(serviceV))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
