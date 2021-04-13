package main

import (
	"context"
	"go-kafka-simple-project/example1"
	"go-kafka-simple-project/example2"
	"log"
)

// kafka settings based on steps 1-5 from https://kafka.apache.org/quickstart

const (
	selectedExample = 1
	topic           = "quickstart-events"
)

func getSupportedBrokerAddresses() []string {
	return []string{"localhost:9092"}
}

func selectExample(e int) func(context.Context, string, []string) {
	log.Printf("Example %v launched.\n", e)
	return [...]func(context.Context, string, []string){example1.Run, example2.Run}[e-1]
}

func main() {
	// produce messages in a new go routine, since both the produce and consume functions are blocking:
	selectExample(selectedExample)(context.Background(), topic, getSupportedBrokerAddresses())
}
