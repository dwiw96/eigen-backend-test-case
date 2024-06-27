package service

import (
	"context"

	books "eigen-backend-test-case/features/books"
)

type booksService struct {
	repo books.RepositoryInterface
	ctx  context.Context
}

func NewBooksService(repo books.RepositoryInterface, ctx context.Context) books.ServiceInterface {
	return &booksService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *booksService) InsertListOfBooks(input []books.Books) (err error) {
	err = s.repo.InsertListOfBooks(input)
	if err != nil {
		return err
	}

	return err
}
