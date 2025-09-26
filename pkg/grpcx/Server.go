package grpcx

import (
	"context"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"strconv"
)

type Server struct {
	*grpc.Server
	Port        int
	RemoteAddr  string
	RemotePort  int
	Addr        string
	Name        string
	cancel      func()
	localserver *consul.AgentService
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func (s *Server) Serve() error {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.Addr = GetOutboundIP()
	port := strconv.Itoa(s.Port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s.localserver = &consul.AgentService{
		Address: s.Addr,
		ID:      s.Name,
		Service: s.Name,
		Port:    s.Port,
	}
	err = s.serverRegister(ctx)
	if err != nil {
		return err
	}
	return s.Server.Serve(l)

}

// 注册服务
func (s *Server) serverRegister(ctx context.Context) error {
	config := &consul.Config{
		Address: fmt.Sprintf("%s:%d", s.RemoteAddr, s.RemotePort),
	}
	client, _ := consul.NewClient(config)
	registration := &consul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s:%d", s.Name, GetOutboundIP(), s.Port),
		Name:    s.Name,
		Address: s.Addr,
		Port:    s.Port,
		Check: &consul.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health/v1/checks/%s", s.RemoteAddr, s.RemotePort, s.Name),
			Interval: "10s",
			Timeout:  "2s",
		},
	}
	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}
	return nil
}

// DeRegister 注销服务
func (s *Server) DeRegister() {

	client, _ := consul.NewClient(&consul.Config{Address: fmt.Sprintf("%s:%d", s.RemoteAddr, s.RemotePort)})
	_, err := client.Catalog().Deregister(&consul.CatalogDeregistration{
		Node:      s.Name,
		Address:   s.Addr,
		ServiceID: fmt.Sprintf("%s-%s:%d", s.Name, GetOutboundIP(), s.Port),
	}, nil)
	if err != nil {
		slog.Error("服务注销失败")
	} else {
		slog.Info("服务注销成功")
	}

}

//func main() {
//	Register()
//	defer DeRegister()
//
//	// 监听端口
//	listen, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		panic(err)
//	}
//	// 创建gprc服务器
//	server := grpc.NewServer(
//		grpc.Creds(insecure.NewCredentials()),
//	)
//	// 注册服务
//	pb.RegisterSayHelloServer(server, &HelloRpc{})
//	log.Println("server running...")
//	// 运行
//	err = server.Serve(listen)
//	if err != nil {
//		panic(err)
//	}
//}
