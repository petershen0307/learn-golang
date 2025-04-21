package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	root "github.com/petershen0307/learn-golang/learn-kafka"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Kafka broker address
	const brokerAddress = "localhost:9092"
	// Kafka topic
	topic := "test-topic"

	// Create a Kafka writer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddress,
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for e := range producer.Events() {
			switch m := e.(type) {
			case *kafka.Message:
				if m.TopicPartition.Error != nil {
					log.Fatalf("Delivery failed: %v", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to %v\n", m.TopicPartition)
				}
			}
		}
		wg.Done()
	}()
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done, exiting...")
			producer.Close()
			wg.Wait()
			return
		default:
			message := root.Message{
				Content:    fmt.Sprintf("Hello, Kafka! %d", i),
				RetryCount: 0,
				Timestamp:  time.Duration(time.Now().UTC().Unix()),
			}
			mdata, err := json.Marshal(message)
			if err != nil {
				log.Fatalf("Error marshalling message: %v", err)
			}

			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          mdata,
			}, nil); err != nil {
				log.Fatalf("Failed to produce message: %s", err)
			}

			fmt.Printf("Sent message: %#v\n", message)
			i++
			time.Sleep(time.Second)
		}
	}
}
