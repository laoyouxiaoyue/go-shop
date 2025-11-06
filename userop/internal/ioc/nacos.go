package ioc

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"shop/userop/internal/config"
)

func InitNacos(c *config.Config) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         c.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            c.NacosConfig.User,
		Password:            c.NacosConfig.Password,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: c.NacosConfig.Host,
			Port:   uint64(c.NacosConfig.Port),
			Scheme: "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: c.NacosConfig.DataID,
		Group:  c.NacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), c)
	if err != nil {
		panic(err)
		return
	}
}
