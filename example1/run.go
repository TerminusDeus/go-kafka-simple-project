package example1

// based on examples from https://www.sohamkamani.com/golang/working-with-kafka/

import "context"

func Run(ctx context.Context, topic string, addresses []string) {
	go produce(ctx, topic, addresses)
	consume(ctx, topic, addresses)
}
