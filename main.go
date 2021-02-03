package main

import (
<<<<<<< Updated upstream
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
=======
	"context"
	"log"
	"test-go/ent"
	"test-go/ent/migrate"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
>>>>>>> Stashed changes
}
