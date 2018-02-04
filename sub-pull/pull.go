package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	name := os.Getenv("PULL_SUBSCRIBER_NAME")
	var mu sync.Mutex
	received := 0

	sub := client.Subscription(name)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		received++
		if received >= 10 {
			cancel()
			msg.Nack() // NACK suggest pub/sub that this client is off connection
			fmt.Print("Subscriber exited. Received more than 10 messages.")
			return
		}
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})

	if err != nil {
		log.Fatal(err)
	}
}
