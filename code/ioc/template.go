package ioc

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop/code/config"
)

func InitYamlTemplateManager() *config.YamlTemplateManager {
	v := viper.New()
	v.SetConfigFile("./templates.yaml") // 直接指定路径
	if err := v.ReadInConfig(); err != nil {
		zap.L().Panic("读取短信模版配置失败")
	}
	ytm := config.YamlTemplateManager{}
	if err := v.Unmarshal(&ytm); err != nil {
		zap.L().Panic("解析短信模版配置失败")
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		v.SetConfigFile("./templates.yaml")
		if err := v.ReadInConfig(); err != nil {
			zap.L().Panic("读取短信模版配置失败")
		}
		err := v.Unmarshal(&ytm)
		if err != nil {
			zap.L().Panic("解析短信模版配置失败")
			return
		}
	})
	return &ytm
}
