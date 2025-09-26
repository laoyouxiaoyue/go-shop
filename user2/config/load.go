package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configLoader interface {
	Load(configPath string) (*Config, error)
}

type ConfigLoaderViper struct {
}

func (c *ConfigLoaderViper) Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./user/config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fal to read config:%w", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fal to unmarshal config:%w", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if cfg.Consul.mode == Remote {

	}
	return &cfg, nil
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config/")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fal to read config:%w", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fal to unmarshal config:%w", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if cfg.Consul.mode == Remote {

	}
	return &cfg, nil
}
