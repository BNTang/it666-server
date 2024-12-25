package config

type Mysql struct {
	// 服务器地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 数据库名
	Dbname string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	// 数据库用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 高级配置
	Config string `mapstructure:"config" json:"config" yaml:"config"`
	// 空闲中的最大连接数
	MaxIdleConns int `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	// Gorm日志等级 "silent"、"error"、"warn"、"info"
	LogMode string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	// 是否写入zap
	LogZap bool `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

// 生成dsn
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
