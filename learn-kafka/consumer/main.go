package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Kafka broker address
	brokerAddress := "localhost:9092"
	// Kafka topic
	topic := "test-topic"

	// Create a Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  brokerAddress,
		"group.id":           "your-group-id", // Set a group ID
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	// Subscribe to the topic
	if err := consumer.Subscribe(topic, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context done, exit reading...")
				return
			default:
				msg, err := consumer.ReadMessage(time.Second) // blocking until a message is received
				if err == nil {
					switch e := err.(type) {
					case kafka.Error:
						if e.IsTimeout() {
							continue
						} else {
							log.Printf("Error reading message: %v", e)
						}
					}
				}
				if msg != nil {
					fmt.Printf("Received message: %s\n", string(msg.Value))
				}
			}
		}
	}()
	<-ctx.Done()
	fmt.Println("Context done, exiting...")
	consumer.Close()
	wg.Wait()
}
