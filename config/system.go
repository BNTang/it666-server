package config

type System struct {
	// 服务端端口号
	Port int `mapstructure:"port" json:"port" yaml:"port"`
	// 密码MD5盐值
	Salt string `mapstructure:"salt" json:"salt" yaml:"salt"`
	// 静态资源路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}
