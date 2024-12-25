package config

type Zap struct {
	// 输出
	Format string `mapstructure:"format" json:"format" yaml:"format"`
	// 日志文件夹
	Director string `mapstructure:"director" json:"director"  yaml:"director"`
	// 是否输出到控制台
	LogInConsole bool `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
