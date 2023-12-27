package main

import (
	"context"
	"re-parteners-tech-task/config"
	"re-parteners-tech-task/handler"
	"re-parteners-tech-task/router"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	handler.Handler(ctx)

	router.Router(ctx, cfg)
}
