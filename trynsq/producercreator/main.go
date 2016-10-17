package main

import (
	"github.com/bitly/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("localhost:4150", config)
	defer w.Stop()

	if err != nil {
		log.Panicf("Failed creating new producer, %v", err)
	}
	qName := "hello_nsq"
	err = w.Publish(qName, []byte("Hello Bilal"))
	if err != nil {
		log.Panicf("Failed publishing %s", qName)
	}
}
