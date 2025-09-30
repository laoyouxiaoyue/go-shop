package grpcx

import (
	"context"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
	"net"
	"strconv"
)

type Server struct {
	*grpc.Server
	Addr        string
	Port        int
	RemoteAddr  string
	RemotePort  int
	Name        string
	tags        []string
	id          string
	cancel      func()
	localserver *consul.AgentService
}

func NewServer(s *grpc.Server, localaddr string, localhost int, address string, port int, name string, tags []string, id string) *Server {
	return &Server{
		Server:     s,
		Addr:       localaddr,
		Port:       localhost,
		RemoteAddr: address,
		RemotePort: port,
		Name:       name,
		tags:       tags,
		id:         id,
	}
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
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.Server, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
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
	err = s.ServerRegister(ctx)
	if err != nil {
		zap.L().Error("consul register failed", zap.Error(err))
		return err
	}
	return s.Server.Serve(l)

}

// 注册服务
func (s *Server) ServerRegister(ctx context.Context) error {
	cfg := consul.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", s.RemoteAddr, s.RemotePort)
	zap.L().Info(cfg.Address)
	client, err := consul.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &consul.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	slog.Info(fmt.Sprintf("%+v", check))
	registration := new(consul.AgentServiceRegistration)
	registration.Name = s.Name
	registration.ID = s.id
	registration.Port = s.Port
	registration.Tags = s.tags
	registration.Address = s.Addr
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
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
