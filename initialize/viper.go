package initialize

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"it666.chat/global"
)

/*
*
Viper
viper是一个很完善的Go项目配置解决方案，很多著名的开源项目都在使用，
比如Hugo,Docker都使用了该库，使用viper可以让我们专注于自己的项目代码，而不用自己写那些配置解析代码

功能：
支持配置key默认值设置
支持读取JSON,TOML,YAML,HCL,envfile和java properties等多种不同类型配置文件
可以监听配置文件的变化，并重新加载配置文件
读取系统环境变量的值
读取存储在远程配置中心的配置数据，如ectd，Consul,firestore等系统，并监听配置的变化
从命令行读取配置
从buffer读取配置
可以显示设置配置的值

例如:
数据库配置
三方登录配置
云存储配置
等等...
*/
func Viper(configType string, path ...string) *viper.Viper {
	var config string
	// 如果没有传入参数，就从环境变量中读取
	if len(path) == 0 {
		// 如果环境变量中没有配置文件路径，就使用默认的
		if configEnv := os.Getenv("IT666_CONFIG"); configEnv == "" {
			config = "config." + configType
			fmt.Printf("正在使用默认的config, 路径为%v\n", config)
		} else {
			// 如果环境变量中有配置文件路径, 就使用环境变量中的
			config = configEnv
			fmt.Printf("正在使用环境变量的config, 路径为%v\n", config)
		}
	} else {
		// 如果有传递参数，就使用传递的参数
		config = path[0]
		fmt.Printf("正在使用func传递的config, 路径为%v\n", config)
	}

	v := viper.New()            // 创建viper实例
	v.SetConfigFile(config)     // 设置配置文件
	v.SetConfigType(configType) // 设置配置文件类型
	err := v.ReadInConfig()     // 如果配置文件不存在，就报错
	if err != nil {
		panic(fmt.Errorf("配置文件错误: %s", err))
	}

	// 将配置文件中的内容读取到结构体中
	if err := v.Unmarshal(&global.IT666_CONFIG); err != nil {
		fmt.Println(err)
	}
	/*
		// 打印从配置文件解析出来的内容
		fmt.Println(global.IT666_CONFIG)
		{
			{127.0.0.1 3306 it666_com root com.it666.www charset=utf8mb4&parseTime=True&loc=Local 10 100 info false}
			{0 127.0.0.1 6379 }
		}
	*/
	// 当配置文件发生变化时，会触发回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生改变:", e.Name)
		// 重新读取配置文件
		if err := v.Unmarshal(&global.IT666_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// Viper 支持让您的应用程序在运行时实时读取配置文件的能力，而无需重新启动应用程序。
	v.WatchConfig() // 监控配置文件变化并热加载程序

	return v
}
