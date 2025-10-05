package config

const (
	Local int = iota
	Remote
)

var Cf *Config

type Config struct {
	Server      ServerConfig `mapstructure:"server" json:"server"`
	Redis       RedisConfig  `mapstructure:"redis" json:"redis"`
	Consul      ConsulConfig `mapstructure:"consul" json:"consul"`
	Smtp        SmtpConfig   `mapstructure:"smtp" json:"smtp"`
	NacosConfig NacosConfig  `mapstructure:"nacos" json:"nacos"`
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

type ServerConfig struct {
	Name          string   `mapstructure:"name" json:"name"`
	Host          string   `mapstructure:"host" json:"host"`
	Port          int      `mapstructure:"port" json:"port"`
	Tags          []string `mapstructure:"tags" json:"tags"`
	Id            string   `mapstructure:"id" json:"id"`
	CheckInterval string   `mapstructure:"check_interval" json:"check_interval"`
	CheckTimeout  string   `mapstructure:"check_timeout" json:"check_timeout"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"pool_size" json:"pool_size"`
}

type ConsulConfig struct {
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	Scheme     string `mapstructure:"scheme" json:"scheme"`
	Token      string `mapstructure:"token" json:"token"`
	Datacenter string `mapstructure:"datacenter" json:"datacenter"`
	WaitTime   string `mapstructure:"wait_time" json:"wait_time"`
}

type SmtpConfig struct {
	From      string   `mapstructure:"from" json:"from"`
	FromAlias []string `mapstructure:"from_alias" json:"from_alias"`
	Host      string   `mapstructure:"host" json:"host"`
	Port      int      `mapstructure:"port" json:"port"`
	Password  string   `mapstructure:"password" json:"password"`
}
