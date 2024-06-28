package service

import (
	"context"
	"errors"
	"fmt"

	books "eigen-backend-test-case/features/books"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type booksService struct {
	repo books.RepositoryInterface
	ctx  context.Context
	dbTx *pgxpool.Pool
}

func NewBooksService(repo books.RepositoryInterface, ctx context.Context, dbTx *pgxpool.Pool) books.ServiceInterface {
	return &booksService{
		repo: repo,
		ctx:  ctx,
		dbTx: dbTx,
	}
}

func (s *booksService) InsertListOfBooks(input []books.Books) (err error) {
	err = s.repo.InsertListOfBooks(input)
	if err != nil {
		return err
	}

	return err
}

func (s *booksService) execTx(ctx context.Context, fn func(*booksService) (bool, error)) (bool, error) {
	tx, err := s.dbTx.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return true, err
	}

	// q := NewDB(tx)
	isServerErr, err := fn(s)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return isServerErr, fmt.Errorf("tx rr = %s, and rb err = %s", err, rbErr)
		}
		return isServerErr, err
	}

	return isServerErr, tx.Commit(ctx)
}

func (s *booksService) BorrowBooks(memberCode, bookCode string) (book books.Books, isServerErr bool, err error) {
	// get the book information
	bookData, err := s.repo.GetBookData(bookCode)
	if err != nil {
		return books.Books{}, true, err
	}

	// get member information
	memberData, err := s.repo.GetMemberData(memberCode)
	if err != nil {
		return books.Books{}, true, err
	}

	// check if the member more that 2 books borrowed
	borroweBookCount, err := s.repo.CheckMemberBorrowedBooks(memberData.ID)
	if err != nil {
		return books.Books{}, true, err
	}
	if borroweBookCount > 1 {
		return books.Books{}, false, errors.New("member borrowed books more than 2")
	}

	// check if member is being penalized
	isPenalized, err := s.repo.CheckIfMemberPenalized(memberData.ID)
	if err != nil {
		return books.Books{}, true, err
	}
	if isPenalized {
		return books.Books{}, false, errors.New("member is pinalized")
	}

	// borrow the book
	isServerErr, err = s.execTx(s.ctx, func(bs *booksService) (bool, error) {
		// check if the book is available or not
		borrowedBook, err := s.repo.CheckIfBookIsAvailable(bookData.ID)
		if err != nil {
			return true, err
		}
		if bookData.Stock-borrowedBook <= 0 {
			return false, errors.New("book is not available")
		}

		err = s.repo.InsertBorrowedBook(bookData.ID, memberData.ID)
		if err != nil {
			return true, err
		}

		return true, err
	})

	return bookData, isServerErr, err
}
