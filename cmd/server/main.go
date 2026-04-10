package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril server...")

	connString := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connString)
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ!")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to open channel: %v\n", err)
		os.Exit(1)
	}

	err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
	if err != nil {
		fmt.Printf("Failed to publish message: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Pause message published!")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("Shutting down...")
}
