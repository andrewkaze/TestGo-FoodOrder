package main

import (
	"bytes"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"


	"time"


	"github.com/streadway/amqp"
)

type Order struct {
	Id      string   `json:"id"`
	IdUser    string `json:"idUser"`
	IdMenu   string `json:"idMenu"`
	status   string `json:"status"`

}

type Delete struct {
	//Id      int   `json:"id"`

	Operation string `json:"opertion"`
}

func (b *Order) TableName() string {
	return "menu"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//var err error
	db, err := gorm.Open(sqlite.Open("C:\\Users\\aandrianto\\TrainingGo\\g2-go\\TheApp\\foodorderdb.db"), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	db.AutoMigrate(Order{})
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {

		for d := range msgs {
			//var user Models.User
			order := Order{}
			log.Printf("Raw message: %s", d.Body)
			json.Unmarshal([]byte(d.Body), &order)
			//switch
			log.Printf("Received a message: %s", order.IdMenu)

				db.Create(order)

			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}