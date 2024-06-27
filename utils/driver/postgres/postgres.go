package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	config "eigen-backend-test-case/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(cfg *config.EnvConfig) *pgxpool.Pool {
	log.Println("<- init db")

	dbAddress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	configPG, err := pgxpool.ParseConfig(dbAddress) // Using environment variables instead of a connection string.
	if err != nil {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.New(context.Background(), configPG.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// defer dbpool.Close()

	return dbpool
}
