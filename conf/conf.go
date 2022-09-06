package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServiceT struct {
	// Name 服务名称
	Name string
	// ServerPort 服务端口
	ServerPort string
	// RunMode 运行模式 debug、test、release
	RunMode string
	// 	AllowOrigin 允许的跨站请求
	AllowOrigin string
}

// ClientT 客户端配置
type ClientT struct {
	// Name 服务名称
	Name string
	// Dsn 地址配置
	Dsn string
	// Timeout 客户端超时
	Timeout int
	// Data
}

// CosT 配置
type CosT struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	S3Region        string `yaml:"s3Region"`
	S3Bucket        string `yaml:"s3Bucket"`
}

// Conf 配置
type ConfT struct {
	// # 服务配置
	Service ServiceT `yaml:"service"`
	// # 客户端配置
	Clients []*ClientT `yaml:"clients"`
	// cos 配置
	Cos CosT `yaml:"cos"`
	// # 客户端配置map
	clientMap map[string]*ClientT `yaml:"clients"`
}

// Conf 全局配置
var Conf = &ConfT{}

// InitConfig 初始化配置文件
func InitConf(configPath string) *ConfT {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// // 将yaml文件解析为一个结构体实例
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})
	Conf.clientMap = make(map[string]*ClientT, 0)
	Conf.initClientMap()
	return Conf
}

// InitClientMap 获取客户端配置
func (c *ConfT) initClientMap() {
	for _, v := range c.Clients {
		c.clientMap[v.Name] = v
	}
}

// GetClient 获取客户端配置
func (c *ConfT) Get(name string) *ClientT {
	val, ok := c.clientMap[name]
	if !ok {
		return nil
	}
	return val
}
