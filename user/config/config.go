package config

import (
	"time"
)

const (
	Local int = iota
	Remote
)

var Cf *Config

type Config struct {
	Server      ServerConfig `mapstructure:"server" json:"server"`
	Mysql       MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Redis       RedisConfig  `mapstructure:"redis" json:"redis"`
	Auth        AuthConfig   `mapstructure:"auth" json:"auth"`
	Consul      ConsulConfig `mapstructure:"consul" json:"consul"`
	NacosConfig NacosConfig  `mapstructure:"nacos" json:"nacos"`
}

type AppConfig struct {
	Env             string        `mapstructure:"env" json:"env"`
	GRPCPort        int           `mapstructure:"grpc_port" json:"grpc_port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout" json:"shutdown_timeout"`
}

type MysqlConfig struct {
	Dsn string `mapstructure:"dsn" json:"dsn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"pool_size" json:"pool_size"`
}

type AuthConfig struct {
	JwtSecret          string `mapstructure:"jwt_secret" json:"jwt_secret"`
	AccessTokenExpire  string `mapstructure:"access_token_expire" json:"access_token_expire"`
	RefreshTokenExpire string `mapstructure:"refresh_token_expire" json:"refresh_token_expire"`
	Issuer             string `mapstructure:"issuer" json:"issuer"`
}

type ServerConfig struct {
	Name          string   `mapstructure:"name" json:"name"`
	Host          string   `mapstructure:"host" json:"host"`
	Port          int      `mapstructure:"port" json:"port"`
	Tags          []string `mapstructure:"tags" json:"tags"`
	Id            string   `mapstructure:"id" json:"id"`
	CheckInterval string   `mapstructure:"check_interval" json:"check_interval"`
	CheckTimeout  string   `mapstructure:"check_timeout" json:"check_timeout"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataID    string `mapstructure:"dataid" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}

type ConsulConfig struct {
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	Scheme     string `mapstructure:"scheme" json:"scheme"`
	Token      string `mapstructure:"token" json:"token"`
	Datacenter string `mapstructure:"datacenter" json:"datacenter"`
	WaitTime   string `mapstructure:"wait_time" json:"wait_time"`
}
