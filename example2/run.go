package example2

// based on examples from https://github.com/segmentio/kafka-go

import "context"

func Run(ctx context.Context, topic string, addresses []string) {
	getTopicsList()
	go produce(ctx, topic, addresses)
	consume(ctx, topic, addresses)
}
