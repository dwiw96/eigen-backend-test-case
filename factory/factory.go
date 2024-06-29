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

	membersDelivery "eigen-backend-test-case/features/members/delivery"
	membersRepository "eigen-backend-test-case/features/members/repository"
	membersService "eigen-backend-test-case/features/members/service"
)

func InitFactory(cfg *config.EnvConfig, db *pgxpool.Pool, router *httprouter.Router, ctx context.Context) {
	log.Println("<- init factory")

	booksRepoInterface := booksRepository.NewBooksRepository(db, ctx)
	booksServiceInterface := booksService.NewBooksService(booksRepoInterface, ctx, db)
	booksDelivery.NewbooksDelivery(router, booksServiceInterface)

	membersRepoInterface := membersRepository.NewMembersRepository(db, ctx)
	membersServiceInterface := membersService.NewMembersService(membersRepoInterface, ctx)
	membersDelivery.NewMembersDelivery(router, membersServiceInterface)

	log.Println("-> init factory")
}
