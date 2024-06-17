package main

import (
	"context"
	"log"

	"langchainGoChatBot/langchain_service"
)

func main() {
	ctx := context.Background()

	go func() {
		if err := langchain_service.Start(ctx); err != nil {
			log.Fatalf("Failed to start langchain service: %v", err)
		}
	}()

	<-ctx.Done()
}
