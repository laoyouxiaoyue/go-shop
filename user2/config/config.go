package config

import "time"

const (
	Local int = iota
	Remote
)

var Cf *Config

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Auth   AuthConfig   `mapstructure:"auth"`
	Consul ConsulConfig `mapstructure:"consul"`
}

type AppConfig struct {
	Env             string        `mapstructure:"env"`
	GRPCPort        int           `mapstructure:"grpc_port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type MysqlConfig struct {
	Dsn string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type AuthConfig struct {
	JwtSecret          string `mapstructure:"jwt_secret"`
	AccessTokenExpire  string `mapstructure:"access_token_expire"`
	RefreshTokenExpire string `mapstructure:"refresh_token_expire"`
	Issuer             string `mapstructure:"issuer"`
}

type ConsulConfig struct {
	mode       int    `mapstructure:"mode"`
	RemoteAddr string `mapstructure:"remote_addr"`
	RemotePort int    `mapstructure:"remote_port"`
	Name       string `mapstructure:"name"`
	Port       int    `mapstructure:"port"`
}
