package ioc

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/zeromicro/go-zero/zrpc"
	"shop/goods_gozero/goodsclient"
	"shop/inventory/inventoryclient"
	"shop/order/internal/config"
	"time"
)

func discoverAddr(c config.Config, serviceName string) (string, error) {
	cfg := consul.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.ConsulConfig.Host, c.ConsulConfig.Port)
	client, err := consul.NewClient(cfg)
	if err != nil {
		return "", err
	}
	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", err
	}
	if len(services) == 0 {
		return "", fmt.Errorf("service %s not found", serviceName)
	}
	svc := services[0]
	return fmt.Sprintf("%s:%d", svc.Service.Address, svc.Service.Port), nil
}

func InitGoodsClient(c config.Config) (goodsclient.Goods, error) {
	addr, err := discoverAddr(c, c.GoodsServiceName)
	if err != nil {
		return nil, err
	}
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{addr},
		NonBlock:  true,
		Timeout:   int64(time.Second * 3),
	})
	return goodsclient.NewGoods(cli), nil
}

func InitInventoryClient(c config.Config) (inventoryclient.Inventory, error) {
	addr, err := discoverAddr(c, c.InventoryServiceName)
	if err != nil {
		return nil, err
	}
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{addr},
		NonBlock:  true,
		Timeout:   int64(time.Second * 3),
	})
	return inventoryclient.NewInventory(cli), nil
}
