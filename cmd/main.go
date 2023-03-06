package main

import (
	"context"
	"log"
	"task2.3.3/internal/config"
	"task2.3.3/internal/logger"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("error with getting config : %v", err)
	}

	ctx := context.Background()

	log, err := logger.Get()
	if err != nil {
		log.Errorf("error in logger.Get: %v", err)
	}

	ctx = context.WithValue(ctx, "log", log)

}
