package utils

import (
	"context"
	app "flavor/internal/app/flavor"
	"fmt"
)

func start() {
	app, _ := app.NewApp(context.Background())
	fmt.Println(app)
}
