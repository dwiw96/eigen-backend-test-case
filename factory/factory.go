package factory

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"

	config "eigen-backend-test-case/config"

	booksDelivery "eigen-backend-test-case/features/books/delivery"
	booksRepository "eigen-backend-test-case/features/books/repository"
	booksService "eigen-backend-test-case/features/books/service"
)

func InitFactory(cfg *config.EnvConfig, db *pgxpool.Pool, router *httprouter.Router, ctx context.Context) {
	log.Println("<- init factory")

	booksRepoInterface := booksRepository.NewBooksRepository(db, ctx)
	booksServiceInterface := booksService.NewBooksService(booksRepoInterface, ctx)
	booksDelivery.NewbooksDelivery(router, booksServiceInterface)

	log.Println("-> init factory")
}
