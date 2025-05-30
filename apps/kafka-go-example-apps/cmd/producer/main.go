package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Extract command-line flags for Kafka broker and topic
	broker := flag.String("broker", "localhost:19092", "Kafka broker address")
	topic := flag.String("topic", "demo-topic", "Kafka topic")
	messages := flag.Int("messages", 100, "Number of messages to produce")
	pause := flag.Int("pause", 1000, "Pause in milliseconds between messages")
	flag.Parse()

	// Validate the input parameters
	fmt.Println("=====================================================")
	fmt.Printf("Kafka Broker                   : %s\n", *broker)
	fmt.Printf("Kafka Topic                    : %s\n", *topic)
	fmt.Printf("Number of messages to produce  : %d\n", *messages)
	fmt.Printf("Pause duration set to          : %d ms\n", *pause)
	fmt.Println("=====================================================")

	// Create a new Kafka writer client to produce messages
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{*broker},
		Topic:   *topic,
	})
	defer writer.Close() // Ensure the writer is closed after use

	fmt.Println("Starting producer...")

	// Produce messages in a loop
	for i := 0; i < *messages; i++ {
		// Create a new Kafka message with a key and value
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(fmt.Sprintf("Hello Message #%d. Timestamp: %s", i, time.Now().Format(time.RFC3339))),
		}

		// Write the message to the Kafka topic
		start := time.Now()
		fmt.Printf("Producing New Message ...\n")
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf("Could not write message. Error: %v", err)
		}
		fmt.Printf("Produced Message: %s\n", msg.Value)
		fmt.Printf("Time taken to produce message: %vms\n", time.Since(start).Milliseconds())

		// Sleep for the specified pause duration before producing the next message
		fmt.Printf("Pausing for %d ms...\n", *pause)
		time.Sleep(time.Duration(*pause) * time.Millisecond)
	}
}
