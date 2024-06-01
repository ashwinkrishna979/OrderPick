package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	controller "OrderPick/controllers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeOrder(orderController *controller.OrderController) {
	fmt.Println("Starting Kafka consumer")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "localhost:34439",

		// Fixed properties
		"group.id":          "kafka-go-getting-started-1",
		"auto.offset.reset": "earliest"})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "purchases"
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatal("Could not subscribe kafka topic:", err)
	}
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			var message map[string]string
			if err := json.Unmarshal(ev.Value, &message); err != nil {
				fmt.Printf("Failed to unmarshal message: %s", err)
				continue
			}
			itemID, exists := message["item_id"]
			if !exists {
				fmt.Println("item_id not found in message")
				continue
			}
			if err := orderController.CreateOrderFromConsumer(itemID); err != nil {
				fmt.Printf("failed to create order: %s", err)
				continue
			}
		}
	}
	c.Close()
}
