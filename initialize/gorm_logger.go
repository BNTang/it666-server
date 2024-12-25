package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"it666.chat/global"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	switch global.IT666_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

type writer struct {
	logger.Writer
}

// NewWriter构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// 重写Printf方法
func (w *writer) Printf(message string, data ...interface{}) {
	if global.IT666_CONFIG.Mysql.LogZap {
		global.IT666_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
