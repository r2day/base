package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	// 设置日志输出到文件
	Logger.SetOutput(os.Stdout)
	// 设置文件名称及行数
	Logger.SetReportCaller(true)
	// 设置日志输出格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志记录级别
	Logger.SetLevel(logrus.DebugLevel)
}
