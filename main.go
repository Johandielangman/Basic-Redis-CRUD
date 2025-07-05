package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Johandielangman/Basic-Redis-CRUD/application"
)

func main() {
	fmt.Println("Starting Server on port 3000")

	app := application.New()

	// ====> ROOT LEVEL CONTEXT
	// If we cancel the root, all the children ones will be cancelled as well.
	// We define a sigint
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
