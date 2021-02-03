package main

import (
	"fmt"
	"test/receive"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := receive.GetConn("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			conn.Publish("test-key", []byte(`{"message":"test"}`))
		}
	}()

	err = conn.StartConsumer("test-queue", "test-key", handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}
	fmt.Println(string(d.Body))
	return true
}
