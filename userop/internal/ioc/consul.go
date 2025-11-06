package ioc

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"shop/userop/internal/config"
)

func ServerRegister(c *config.Config) {
	cfg := consul.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.ConsulConfig.Host, c.ConsulConfig.Port)
	zap.L().Info(cfg.Address)
	client, err := consul.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &consul.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", c.ServerConfig.Host, c.ServerConfig.Port),
		Timeout:                        c.ConsulConfig.WaitTime,
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(consul.AgentServiceRegistration)
	registration.Name = c.ServerConfig.Name
	registration.ID = c.ServerConfig.Id
	registration.Port = c.ServerConfig.Port
	registration.Tags = c.ServerConfig.Tags
	registration.Address = c.ServerConfig.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return
}
