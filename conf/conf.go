package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ServiceT 服务配置
type ServiceT struct {
	// Name 服务名称
	Name string
	// ServerPort 服务端口
	ServerPort string
	// RunMode 运行模式 debug、test、release
	RunMode string
	// 	AllowOrigin 允许的跨站请求
	AllowOrigin string
	// 	EndpointPrefix api前缀
	EndpointPrefix string
	// 	SecretKey api 加密/登录加密
	SecretKey string
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
	Endpoint        string `yaml:"endpoint"`
}

// CloudT 配置
type CloudT struct {
	SecretID  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
	// # 客户端配置
	Sms []*SmsT `yaml:"sms"`

	// # 短信配置map
	smsMap map[string]*SmsT `yaml:"sms_map"`
}

// SmsT 短信配置
type SmsT struct {
	Name string `yaml:"name"`
	// 短信应用id
	AppID string `yaml:"appId"`
	// 签名-名称
	SignName string `yaml:"signName"`
	// 模版id
	TplId string `yaml:"tplId"`
}

// Configuration 配置
type Configuration struct {
	// # 服务配置
	Service ServiceT `yaml:"service"`
	// # 客户端配置
	Clients []*ClientT `yaml:"clients"`
	// Cloud 云配置
	Cloud CloudT `yaml:"cloud"`
	// cos 配置
	Cos CosT `yaml:"cos"`

	// # 客户端配置map
	clientMap map[string]*ClientT `yaml:"clients_map"`
}

// Conf 全局配置
var Conf = &Configuration{}

// InitConf 初始化配置文件
func InitConf(configPath string) *Configuration {
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

	// 将配置名称映射到配置上
	Conf.clientMap = make(map[string]*ClientT, 0)
	Conf.initClientMap()

	Conf.Cloud.smsMap = make(map[string]*SmsT, 0)
	Conf.initSmsMap()

	return Conf
}

// InitClientMap 获取客户端配置
func (c *Configuration) initClientMap() {
	for _, v := range c.Clients {
		c.clientMap[v.Name] = v
	}
}

// initSmsMap 初始化短信配置
func (c *Configuration) initSmsMap() {
	for _, v := range c.Cloud.Sms {
		c.Cloud.smsMap[v.Name] = v
	}
}

// Get 获取客户端配置
func (c *Configuration) Get(name string) *ClientT {
	val, ok := c.clientMap[name]
	if !ok {
		return nil
	}
	return val
}

// GetSms 获取短信配置
func (c *Configuration) GetSms(name string) *SmsT {
	val, ok := c.Cloud.smsMap[name]
	if !ok {
		return nil
	}
	return val
}
