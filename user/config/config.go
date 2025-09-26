package config

import (
	"go.uber.org/zap"
	"time"
)

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

func PrintConfig() {
	if Cf == nil {
		zap.S().Warn("配置未初始化")
		return
	}

	zap.S().Info("=== 应用配置信息 ===")

	// 应用配置
	zap.S().Infow("应用配置",
		"env", Cf.App.Env,
		"grpc_port", Cf.App.GRPCPort,
		"shutdown_timeout", Cf.App.ShutdownTimeout,
	)

	// MySQL 配置（敏感信息脱敏）
	mysqlDSN := Cf.Mysql.Dsn
	if len(mysqlDSN) > 20 { // 简单脱敏，只显示部分信息
		mysqlDSN = mysqlDSN[:20] + "..."
	}
	zap.S().Infow("MySQL配置",
		"dsn_length", len(Cf.Mysql.Dsn),
		"dsn_preview", mysqlDSN,
	)

	// Redis 配置
	zap.S().Infow("Redis配置",
		"host", Cf.Redis.Host,
		"port", Cf.Redis.Port,
		"password_set", Cf.Redis.Password != "",
		"db", Cf.Redis.DB,
		"pool_size", Cf.Redis.PoolSize,
	)

	// 认证配置（敏感信息脱敏）
	jwtSecret := Cf.Auth.JwtSecret
	if len(jwtSecret) > 5 {
		jwtSecret = jwtSecret[:5] + "..."
	}
	zap.S().Infow("认证配置",
		"jwt_secret_length", len(Cf.Auth.JwtSecret),
		"jwt_secret_preview", jwtSecret,
		"access_token_expire", Cf.Auth.AccessTokenExpire,
		"refresh_token_expire", Cf.Auth.RefreshTokenExpire,
		"issuer", Cf.Auth.Issuer,
	)

	// Consul 配置
	zap.S().Infow("Consul配置",
		"remote_addr", Cf.Consul.RemoteAddr,
		"remote_port", Cf.Consul.RemotePort,
		"name", Cf.Consul.Name,
		"port", Cf.Consul.Port,
	)
}
