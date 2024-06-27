package main

import (
	"context"
	"log"
	"time"

	api "eigen-backend-test-case/api"
	config "eigen-backend-test-case/config"
	factory "eigen-backend-test-case/factory"
	postgres "eigen-backend-test-case/utils/driver/postgres"
)

func main() {
	log.Println("eigen-book")
	cfg := config.GetConfig()
	dbPool := postgres.InitDB(cfg)
	defer dbPool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	router := api.SetUpRouter()
	factory.InitFactory(cfg, dbPool, router, ctx)

	api.StartServer(cfg.SERVER_PORT, router)
}
