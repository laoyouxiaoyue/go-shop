package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop/web/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./web/config-debug.yaml")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&global.ServerConfig)
	if err != nil {
		panic(err)
		return
	}
	zap.S().Infof("server config: %v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		zap.S().Infof("server config changed %s", e.Name)
		_ = v.ReadInConfig()
		err := v.Unmarshal(&global.ServerConfig)
		if err != nil {
			panic(err)
			return
		}
		zap.S().Infof("server config: %v", global.ServerConfig)
	})
}
