package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"shop/web/global"
)

// DiscoverAddr 通过 Consul 健康检查发现服务可用地址
func DiscoverAddr(serviceName string) (string, error) {
	cfg := api.DefaultConfig()
	// 从配置读取 Consul 地址
	host := global.ServerConfig.Consul.Host
	port := global.ServerConfig.Consul.Port
	if host == "" || port == 0 {
		host = "127.0.0.1"
		port = 8500
	}
	cfg.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return "", err
	}
	// 只取第一个健康实例
	svcs, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", err
	}
	if len(svcs) == 0 {
		return "", fmt.Errorf("service %s not found", serviceName)
	}
	inst := svcs[0]
	return fmt.Sprintf("%s:%d", inst.Service.Address, inst.Service.Port), nil
}

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		HTTP:                           "http://127.0.0.1:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return nil
}
func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
		return
	}
	for key, _ := range services {
		fmt.Println(key)
	}
}
func FilterService() {
	//cfg := api.DefaultConfig()
	//cfg.Address = "127.0.0.1:8500"
	//
	//client, err := api.NewClient(cfg)
	//if err != nil {
	//	panic(err)
	//}
	//client.Agent().ServicesWithFilter()
}
func main() {
	//AllService()
	_ = Register("127.0.0.1", 8021, "test", []string{"123"}, "123")
}
