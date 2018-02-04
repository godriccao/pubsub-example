package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	// Sets the name for the new topic.
	topic := os.Getenv("TOPIC")
	msg := "蒙蒙下班了？🐖回家晚点，晚上吃草"

	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("failed to get result id: %v", err)
	}

	fmt.Printf("Published a message; msg ID: %v; msg: %s", id, msg)
}
