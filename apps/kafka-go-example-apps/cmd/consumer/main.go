package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Extract command-line flags for Kafka broker, topic, and consumer group
	broker := flag.String("broker", "localhost:19092", "Kafka broker address")
	topic := flag.String("topic", "demo-topic", "Kafka topic")
	group := flag.String("group", "demo-group", "Kafka consumer group")
	flag.Parse()

	// Validate the input parameters
	fmt.Println("=====================================================")
	fmt.Printf("Kafka Broker   : %s\n", *broker)
	fmt.Printf("Kafka Topic    : %s\n", *topic)
	fmt.Printf("Consumer Group : %s\n", *group)
	fmt.Println("=====================================================")

	// Create a new Kafka reader client to consume messages
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{*broker},
		GroupID: *group,
		Topic:   *topic,
	})
	defer reader.Close() // Ensure the reader is closed after use

	fmt.Println("Starting consumer. It will run indefinitely, until stopped...")

	// Consume messages in a loop
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Could not read message. Error: %v", err)
		}
		fmt.Printf("Consumed: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}
