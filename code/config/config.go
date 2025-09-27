package config

import "time"

const (
	Local int = iota
	Remote
)

var Cf *Config

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Consul ConsulConfig `mapstructure:"consul"`
	Smtp   SmtpConfig   `mapstructure:"smtp"`
}
type AppConfig struct {
	Env             string        `mapstructure:"env"`
	GRPCPort        int           `mapstructure:"grpc_port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type ConsulConfig struct {
	mode       int    `mapstructure:"mode"`
	RemoteAddr string `mapstructure:"remote_addr"`
	RemotePort int    `mapstructure:"remote_port"`
	Name       string `mapstructure:"name"`
	Port       int    `mapstructure:"port"`
}

type SmtpConfig struct {
	From      string   `mapstructure:"from"`
	FromAlias []string `mapstructure:"from_alias"`
	Host      string   `mapstructure:"host"`
	Port      int      `mapstructure:"port"`
	Password  string   `mapstructure:"password"`
}
