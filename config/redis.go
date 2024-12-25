package config

type Redis struct {
	// redis的哪个数据库
	DB int `mapstructure:"db" json:"db" yaml:"db"`
	// 地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
