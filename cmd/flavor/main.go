package main

import (
	"context"
	app "flavor/internal/app/flavor"
	"log"
)

func main() {
	ctx := context.Background()

	_, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

}
