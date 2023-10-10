package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		fmt.Println("Error creating Pulsar client:", err)
		return
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "account-create-topic",
		SubscriptionName: "my-subscription",
		Type:             pulsar.Shared,
	})
	if err != nil {
		fmt.Println("Error creating Pulsar consumer:", err)
		return
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()

	fmt.Println("Pulsar consumer started. Waiting for messages...")

	for {
		select {
		case msg := <-consumer.Chan():
			// 处理接收到的消息
			fmt.Printf("Received message: %s\n", string(msg.Payload()))
			consumer.Ack(msg)
		case <-ctx.Done():
			return
		case <-signals:
			return
		}
	}
}
