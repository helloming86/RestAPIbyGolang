package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// 结构体
type Config struct {
	Name string
}
// 结构体方法initConfig
// 设置并解析配置文件
// 根据 initConfig 的 定义，先从配置文件读取配置；如果有设置环境变量，再从环境变量读取配置
func (c *Config) initConfig() error  {
	// 1. 读配置文件
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	// 2. 读环境变量
	viper.AutomaticEnv() // 读取匹配的环境变量
	viper.SetEnvPrefix("MIAPI") // 读取环境变量的前缀为 MIAPI
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	//  viper.ReadInConfig() 函数最终会调用 Viper 解析配置文件。
	if err := viper.ReadInConfig(); err !=nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}


func Init(cfg string) error {
	// 简短声明结构体变量
	c := Config{Name: cfg}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}