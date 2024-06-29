package repository

import (
	"context"
	books "eigen-backend-test-case/features/books"
	"fmt"
	"log"
	"strings"
	"time"

	helper "eigen-backend-test-case/utils/helper"

	"github.com/jackc/pgx/v5"
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
	query := "INSERT INTO books(code, title, author, stock, total_amount) VALUES"
	var placeholder []string
	var arguments []interface{}

	for i, v := range input {
		params := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", (i*5)+1, (i*5)+2, (i*5)+3, (i*5)+4, (i*5)+5)
		placeholder = append(placeholder, params)
		arguments = append(arguments, v.Code, v.Title, v.Author, v.Stock, v.Stock)
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

func (r *booksRepository) GetBookData(bookCode string) (book books.Books, err error) {
	query := `SELECT * FROM books WHERE code=$1`

	row := r.db.QueryRow(r.ctx, query, bookCode)
	err = row.Scan(&book.ID, &book.Code, &book.Title, &book.Author, &book.Stock, &book.TotalAmount)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("error get book data for code: %s, err: %v\n", bookCode, err)
			return books.Books{}, err
		}
		errMsg := fmt.Errorf("error scan get book data")
		log.Printf("%v, err: %v\n", errMsg, err)
		return books.Books{}, errMsg
	}

	return
}

func (r *booksRepository) CheckIfBookIsAvailable(bookID int) (res int, err error) {
	query := `SELECT COUNT(*) FROM borrowed_books WHERE book_id=$1 AND is_returned=FALSE`

	err = r.db.QueryRow(r.ctx, query, bookID).Scan(&res)
	if err != nil {
		errMsg := fmt.Errorf("error check if book is available")
		log.Printf("%v, err: %v\n", errMsg, err)
		return 0, errMsg
	}

	return
}

func (r *booksRepository) GetMemberData(memberCode string) (member books.Member, err error) {
	query := `SELECT * FROM members WHERE code = $1`

	err = r.db.QueryRow(r.ctx, query, memberCode).Scan(&member.ID, &member.Code, &member.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("error get member data for code: %s, err: %v\n", memberCode, err)
			return books.Member{}, err
		}
		errMsg := fmt.Errorf("error scan get member data")
		log.Printf("%v, err: %v\n", errMsg, err)
		return books.Member{}, errMsg
	}

	return
}

func (r *booksRepository) CheckMemberBorrowedBooks(memberID int) (res int, err error) {
	query := `SELECT COUNT(*) FROM borrowed_books WHERE member_id = $1 AND is_returned=FALSE `

	err = r.db.QueryRow(r.ctx, query, memberID).Scan(&res)
	if err != nil {
		errMsg := fmt.Errorf("error check member borrowed books")
		log.Printf("%v, err: %v\n", errMsg, err)
		return 0, errMsg
	}

	return
}

func (r *booksRepository) CheckIfMemberPenalized(memberID int) (res bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM penalized_members WHERE member_id = $1 AND penalty_end > NOW())`

	err = r.db.QueryRow(r.ctx, query, memberID).Scan(&res)
	if err != nil {
		errMsg := fmt.Errorf("error check if member penalized")
		log.Printf("%v, err: %v\n", errMsg, err)
	}

	return
}

func (r *booksRepository) InsertBorrowedBook(bookID, memberID int) (err error) {
	query := `INSERT INTO borrowed_books(book_id, member_id, borrowed_at, is_returned) 
	VALUES($1, $2, $3, $4)`

	res, err := r.db.Exec(r.ctx, query, bookID, memberID, helper.FormatGoTime(time.Now().UTC()), false)
	if err != nil {
		errMsg := fmt.Errorf("error insert borrowed book")
		log.Printf("%v, err:%v\n", errMsg, err)
		return errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no books added to database for insert borrowed book")
		log.Println("no rows are affected for insert borrowed book")
		return errMsg
	}

	return err
}

func (r *booksRepository) CheckMemberBorrowedValidBook(memberID, bookID int) (res bool, err error) {
	query := `SELECT EXISTS(SELECT * FROM borrowed_books WHERE member_id = $1 AND book_id = $2 AND is_returned = FALSE)`

	err = r.db.QueryRow(r.ctx, query, memberID, bookID).Scan(&res)
	if err != nil {
		errMsg := fmt.Errorf("error check member borrowed valid book")
		log.Printf("%v, err: %v\n", errMsg, err)
		return false, errMsg
	}

	return
}

func (r *booksRepository) UpdateBorrowedBookToReturned(id int) (returnedTime time.Time, err error) {
	query := `UPDATE borrowed_books SET returned_at = $1, is_returned = TRUE WHERE id = $2`

	returnedTime = helper.FormatGoTime(time.Now()).UTC()
	res, err := r.db.Exec(r.ctx, query, returnedTime, id)
	if err != nil {
		errMsg := fmt.Errorf("error update borrowed book to returned")
		log.Printf("%v, err:%v\n", errMsg, err)
		return time.Time{}, errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no row updated for update borrowed book to returned")
		log.Println("no rows are affected for update borrowed book to returned")
		return time.Time{}, errMsg
	}

	return returnedTime, err
}

func (r *booksRepository) GetBorrowedBookData(memberID, bookID int) (res books.BorrowedBooks, err error) {
	query := `SELECT * FROM borrowed_books WHERE member_id = $1 AND book_id = $2 ORDER BY id DESC LIMIT 1`

	err = r.db.QueryRow(r.ctx, query, memberID, bookID).Scan(&res.ID, &res.BookID, &res.MemberID, &res.BorrowedAt, &res.ReturnedAt, &res.IsReturned)
	if err != nil {
		errMsg := fmt.Errorf("error get borrowed book data")
		log.Printf("%v, err: %v\n", errMsg, err)
		return books.BorrowedBooks{}, errMsg
	}

	return
}

func (r *booksRepository) InsertPenalty(memberID int, pinaltyStart, pinaltyEnd time.Time) (err error) {
	query := `INSERT INTO penalized_members(member_id, penalty_start, penalty_end) 
	VALUES($1, $2, $3)`

	res, err := r.db.Exec(r.ctx, query, memberID, pinaltyStart, pinaltyEnd)
	if err != nil {
		errMsg := fmt.Errorf("error insert penalty")
		log.Printf("%v, err:%v\n", errMsg, err)
		return errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no row affected for insert penalty")
		log.Println("no rows are affected for insert penalty")
		return errMsg
	}

	return
}

func (r *booksRepository) UpdateBookStock(bookID, amount int) (err error) {
	query := `UPDATE books SET stock = stock + $1 WHERE id = $2`

	res, err := r.db.Exec(r.ctx, query, amount, bookID)
	if err != nil {
		errMsg := fmt.Errorf("error update book stock")
		log.Printf("%v, err:%v\n", errMsg, err)
		return errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no row updated for update book stock")
		log.Println("no rows are affected for update book stock")
		return errMsg
	}

	return
}

func (r *booksRepository) ListExistingBooks() (res []books.Books, err error) {
	query := `SELECT * FROM books WHERE stock > 0 ORDER BY title ASC`

	rows, err := r.db.Query(r.ctx, query)
	if err != nil {
		errMsg := fmt.Errorf("error get existing books")
		log.Printf("%v, err: %v\n", errMsg, err)
		return nil, errMsg
	}
	defer rows.Close()

	for rows.Next() {
		var temp books.Books
		err := rows.Scan(&temp.ID, &temp.Code, &temp.Title, &temp.Author, &temp.Stock, &temp.TotalAmount)
		if err != nil {
			errMsg := fmt.Errorf("error scanning get existing books")
			log.Printf("%v, err: %v\n", errMsg, err)
			return nil, errMsg
		}

		res = append(res, temp)
	}

	return
}
