package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	brokerAddress := "localhost:9092"
	// Kafka topic
	topic := "test-topic"

	// Create a Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// GroupID:        uuid.New().String(),
		GroupBalancers: []kafka.GroupBalancer{kafka.RangeGroupBalancer{}},
		StartOffset:    kafka.FirstOffset,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
	})
	defer reader.Close()

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done, exiting...")
			return
		default:
			m, err := reader.FetchMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			fmt.Printf("Received message: %s\n", string(m.Value))
			time.Sleep(time.Second)
		}
	}
}
