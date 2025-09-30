package config

const (
	Local int = iota
	Remote
)

var Cf *Config

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Consul ConsulConfig `mapstructure:"consul"`
	Smtp   SmtpConfig   `mapstructure:"smtp"`
}
type ServerConfig struct {
	Name          string   `mapstructure:"name"`
	Host          string   `mapstructure:"host"`
	Port          int      `mapstructure:"port"`
	Tags          []string `mapstructure:"tags"`
	Id            string   `mapstructure:"id"`
	CheckInterval string   `mapstructure:"check_interval"`
	CheckTimeout  string   `mapstructure:"check_timeout"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type ConsulConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Scheme     string `mapstructure:"scheme"`
	Token      string `mapstructure:"token"`
	Datacenter string `mapstructure:"datacenter"`
	WaitTime   string `mapstructure:"wait_time"`
}

type SmtpConfig struct {
	From      string   `mapstructure:"from"`
	FromAlias []string `mapstructure:"from_alias"`
	Host      string   `mapstructure:"host"`
	Port      int      `mapstructure:"port"`
	Password  string   `mapstructure:"password"`
}
