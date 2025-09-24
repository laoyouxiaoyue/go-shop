package config

type UserServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
type ServerConfig struct {
	Name       string           `mapstructure:"name"`
	Port       string           `mapstructure:"port"`
	UserServer UserServerConfig `mapstructure:"user_server"`
	Jwt        JWTConfig        `mapstructure:"jwt"`
}
