package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"it666.com/utils"
)

/**
* 1.什么是zap
* zap是Uber开源的一个高性能的日志库，它的性能比logrus和log4go都要好，它的性能是logrus的10倍，log4go的100倍。
* 2.为什么使用日志库
* 日志库的作用是记录程序运行过程中的信息，比如：程序启动时间、运行时间、运行状态、运行错误等等。
* 项目上线之后难免会遇到BUG出现的情况，日志可以记录这些BUG出现的地点从而方便进行快速定位和排查
* 3.zap基本使用
* 去看文档丫, 去看文档丫.
* 4.不想看怎么办?
* 你可以直接拷贝我写的下列代码, 你可以直接拿来用.
 */
func Zap(logDir string, isConsole bool, format string) (logger *zap.Logger) {
	// 判断日志文件夹是否存在，不存在则创建
	if ok, _ := utils.PathExists(logDir); !ok {
		fmt.Printf("创建 %v 目录\n", logDir)
		_ = os.Mkdir(logDir, os.ModePerm)
	}

	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	// 构造日志
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", logDir), debugPriority, isConsole, format),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", logDir), infoPriority, isConsole, format),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", logDir), warnPriority, isConsole, format),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", logDir), errorPriority, isConsole, format),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler, isConsole bool, format string) (core zapcore.Core) {
	writer := getWriteSyncer(fileName, true)
	if isConsole {
		return zapcore.NewCore(getEncoder(format), writer, level)
	}
	return zapcore.NewCore(getEncoder(format), writer, level)
}

// 自定义日志分割
func getWriteSyncer(file string, isConsole bool) zapcore.WriteSyncer {
	// 借助 lumberjack库 协助完成日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	// 判断是否需要 同时输出到控制台和文件
	if isConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 编码器(如何写入日志, 是以控制台格式还是json格式)
func getEncoder(format string) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// 获取自定义日志配置
func getEncoderConfig() (config zapcore.EncoderConfig) {
	// 日志格式配置
	config = zapcore.EncoderConfig{
		MessageKey:     "message",                      // 日志内容对应的key名，此参数必须不为空，否则日志主体不处理
		LevelKey:       "level",                        // 日志级别对应的key名
		TimeKey:        "time",                         // 时间对应的key名
		NameKey:        "logger",                       // logger名对应的key名
		CallerKey:      "caller",                       // 调用者对应的key名
		StacktraceKey:  "stacktrace",                   // 栈追踪的key名
		LineEnding:     zapcore.DefaultLineEnding,      // 默认换行符"\n"
		EncodeLevel:    customEncodeLevel,              // 日志等级序列为小写字符串，如:InfoLevel被序列化为 "info"
		EncodeTime:     customTimeEncoder,              // 日志时间格式显示
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间序列化，Duration为经过的浮点秒数
		EncodeCaller:   customEncodeCaller,             // 日志行号显示
	}
	return config
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006/01/02 - 15:04:05.000") + "]")
}

// 自定义日志级别显示
func customEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// 自定义行号显示
func customEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
