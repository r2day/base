package consumer

import (
	logger "github.com/r2day/base/log"
)

// CT 消费者定义
type CT struct {
	ProcessFunc map[string]func([]byte) []byte
}

func NewConsumer() *CT {
	return &CT{
		ProcessFunc: make(map[string]func([]byte) []byte, 0),
	}
}

func (p *CT) Register(appId string, f func([]byte) []byte) {
	p.ProcessFunc[appId] = f
	logger.Logger.Println("register new appId successful")

}

// Consumer 可以直接在xxx-consumer 类型使用
var Consumer = NewConsumer()

// func init() {
//	Consumer.Register(enum.QueueAppIdPlace, orderPlace)
//	Consumer.Register(enum.QueueAppIdOutOfPayTime, outOfPayTime)
//	Consumer.Register(enum.QueueAppIdPayDone, orderPayDone)
//}

// in main.go
// 	// 消息
//	mq := middle.InitAMQP(mqConf.Dsn, enum.QueueOrder)
//	mq.Receive(consumer.Handler)

// Handler 消费者
func Handler(appId string, messageId string, payload []byte) []byte {
	logger.Logger.Infof("appId: %s | messageId: %s", appId, messageId)

	f, ok := Consumer.ProcessFunc[appId]
	if !ok {
		logger.Logger.Warnf("no found any process function for appId: %s", appId)
		return nil
	}
	result := f(payload)
	logger.Logger.Info("the result of handler: %s", result)
	return nil
}
