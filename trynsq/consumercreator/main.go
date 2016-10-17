package main

import (
	"sync"
	"github.com/bitly/go-nsq"
	"log"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("hello_nsq", "ch", config)
	if err != nil {
		log.Panicf("Failed creating new new consumer %v", err)
	}
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("The message", message)
		theMessage := string(message.Body[:])
		log.Printf("The body%s", theMessage)
		wg.Done()
		return nil
	}))

	err = q.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Panicf("Failed connect to nsqlookupd, %v", err)
	}

	wg.Wait()
}
