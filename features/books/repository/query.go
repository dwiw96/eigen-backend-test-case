package repository

import (
	"context"
	books "eigen-backend-test-case/features/books"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type booksRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewBooksRepository(db *pgxpool.Pool, ctx context.Context) books.RepositoryInterface {
	return &booksRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *booksRepository) InsertListOfBooks(input []books.Books) (err error) {
	query := "INSERT INTO books(code, title, author, stock) VALUES"
	var placeholder []string
	var arguments []interface{}

	for i, v := range input {
		params := fmt.Sprintf("($%d, $%d, $%d, $%d)", (i*4)+1, (i*4)+2, (i*4)+3, (i*4)+4)
		placeholder = append(placeholder, params)
		arguments = append(arguments, v.Code, v.Title, v.Author, v.Stock)
	}

	queryPlaceholder := strings.Join(placeholder, ",")
	query += queryPlaceholder

	res, err := r.db.Exec(r.ctx, query, arguments...)
	if err != nil {
		errMsg := fmt.Errorf("error insert list of books")
		log.Printf("%v, err:%v\n", errMsg, err)
		return errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no books added to database for insert list of books")
		log.Println("no rows are affected for insert list of books")
		return errMsg
	}

	return err
}
