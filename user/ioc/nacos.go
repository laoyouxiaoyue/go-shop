package ioc

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"shop/user/config"
)

func InitNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Cf.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            config.Cf.NacosConfig.User,
		Password:            config.Cf.NacosConfig.Password,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.Cf.NacosConfig.Host,         // 只需要 IP 或域名，不要带端口
			Port:   uint64(config.Cf.NacosConfig.Port), // 端口单独设置
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
		DataId: config.Cf.NacosConfig.DataID,
		Group:  config.Cf.NacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &config.Cf)
	if err != nil {
		panic(err)
		return
	}
}
