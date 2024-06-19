package main

import (
	"context"
	"langchainGoChatBot/langchain_service"
	"log"
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
