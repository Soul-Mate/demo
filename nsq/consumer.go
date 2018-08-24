package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"fmt"
)

type data struct {
}

func (d *data) HandleMessage(message *nsq.Message) error {
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("write_test", "ch", config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(message)
		return nil
	}))
	<-consumer.StopChan
}
