package main

import (
	"chi2/application"
	"context"
	"fmt"
)

func main() {
	fmt.Println("Starting Server on port 3000")

	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
