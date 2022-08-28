package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// DBConfig DBConfigDSN 数据库配置
type DBConfig struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Charset    string
}

// WxConf 微信配置
type WxConf struct {
	AppId           string
	Code2SessionUrl string
	Secret          string
	GrantType       string
}

// Config 总配置
type Config struct {
	ServerName  string
	ServerPort  string
	RunMode     string
	SecretKey   string
	AllowOrigin string
	Wx          WxConf
	ConfDB      DBConfig
	DataDB      DBConfig
	DataDSN     string
	// 请求参数签名 (配置在前端所在服务器上TODO)
	SignKey string
	// amqp地址
	MQConf string
	// ES 地址
	ESConf string
}

// ConfInstance xx
var ConfInstance = &Config{}

// InitConfig 初始化配置文件
func InitConfig(configPath string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// // 将yaml文件解析为一个结构体实例
	err = viper.Unmarshal(&ConfInstance)
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&ConfInstance)
		if err != nil {
			panic(err)
		}
	})
	return ConfInstance
}
