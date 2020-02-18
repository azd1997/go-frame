package nsq

import (
	"fmt"
	"github.com/azd1997/go-frame/config"
	"github.com/azd1997/go-frame/logger"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"os"
)

func StartNsqServer(conf config.NsqConfig) {
	cfg := nsq.NewConfig()
	q, _ := nsq.NewConsumer(conf.Topic, conf.Channel, cfg)
	q.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		// 接收到消息时做的事

		// 打印
		fmt.Println(string(msg.Body))
		// 日志
		logger.Logger().Info("receive", zap.String(conf.Topic, string(msg.Body)))

		return nil
	}))

	//err := q.ConnectToNSQLookupd(conf.NsqLookupdAddr)
	err := q.ConnectToNSQD(conf.NsqdAddr)
	if err != nil {
		logger.Logger().Error("connect to nsqlookupd failed: ", zap.Error(err))
		os.Exit(-1)
	}
}
