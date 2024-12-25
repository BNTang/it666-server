package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`    // mysql配置
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`    // redis配置
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`          // zap配置
	System System `mapstructure:"system" json:"system" yaml:"system"` // 系统配置
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`       // 跨域配置
}
