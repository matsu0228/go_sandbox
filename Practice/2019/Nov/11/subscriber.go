package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func run() error {
	ctx := context.Background()
	c, err := newClient(ctx)
	if err != nil {
		return err
	}

	subName := "test-scriber"
	sub := c.Subscription(subName)
	err = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
