package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

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
