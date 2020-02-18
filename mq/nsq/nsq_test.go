package nsq

import (
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
)

func TestProducer(t *testing.T)  {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)

	if err != nil {
		log.Panic(err)
	}

	err = p.Publish("eiger", []byte("hello new msg222"))
	if err != nil {
		log.Panic(err)
	}
}
