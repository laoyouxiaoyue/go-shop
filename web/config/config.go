package config

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
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
	Name        string           `mapstructure:"name" json:"name"`
	Port        string           `mapstructure:"port" json:"port"`
	UserServer  UserServerConfig `mapstructure:"user_server" json:"user_server"`
	Jwt         JWTConfig        `mapstructure:"jwt" json:"jwt"`
	NacosConfig NacosConfig      `mapstructure:"nacos" json:"nacos"`
}
