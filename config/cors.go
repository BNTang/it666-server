package config

type CORS struct {
	// 匹配模式: allow-all放行所有/strict-whitelist严格白名单/whitelist白名单
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
	// 名单列表
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	// 允许的域名
	AllowOrigin string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
	// 允许的方法
	AllowMethods string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	// 允许的头部
	AllowHeaders string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	// 允许暴露的头部
	ExposeHeaders string `mapstructure:"expose-headers" json:"expose-headers" yaml:"expose-headers"`
	// 允许携带凭证
	AllowCredentials bool `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
}
