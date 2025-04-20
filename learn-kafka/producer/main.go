package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	root "github.com/petershen0307/learn-golang/learn-kafka"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	const brokerAddress = "localhost:9092"
	// Kafka topic
	const topic = "test-topic"

	// Create a Kafka writer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.RoundRobin{},
	}
	defer writer.Close()
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done, exiting...")
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

			if err := writer.WriteMessages(ctx,
				kafka.Message{
					Value: mdata,
				}); err != nil {
				log.Fatalf("Error writing message: %v", err)
			}
			fmt.Printf("Sent message: %#v\n", message)
			i++
			time.Sleep(time.Second)
		}
	}
}
