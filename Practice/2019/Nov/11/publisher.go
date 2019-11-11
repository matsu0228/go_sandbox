package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func run() error {
	ctx := context.Background()
	c, err := newClient(ctx)
	if err != nil {
		return err
	}
	topic := c.Topic(topicName)
	defer topic.Stop()

	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	results = append(results, r)
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
