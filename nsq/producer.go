package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	)

func main() {
	var err error
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Publish("write_test", []byte("publish_test_case"))
	if err != nil {
		log.Fatalf("should lazily connect - %s", err)
	}
}
