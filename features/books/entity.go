package books

import (
	"database/sql"
)

type Books struct {
	ID     int
	Code   string
	Title  string
	Author string
	Stock  int
}

type Member struct {
	ID   int
	Code string
	Name string
}

type BorrowedBooks struct {
	ID         int
	BookID     int
	MemberID   int
	BorrowedAt sql.NullTime
	ReturnedAt sql.NullTime
	IsReturned bool
}

type RepositoryInterface interface {
	InsertListOfBooks(input []Books) (err error)
	GetBookData(bookCode string) (book Books, err error)
	CheckIfBookIsAvailable(bookID int) (res int, err error)
	GetMemberData(memberCode string) (member Member, err error)
	CheckMemberBorrowedBooks(memberID int) (res int, err error)
	CheckIfMemberPenalized(memberID int) (res bool, err error)
	InsertBorrowedBook(bookID, memberID int) (err error)
	CheckMemberBorrowedValidBook(memberID, bookID int) (res bool, err error)
	UpdateBorrowedBookToReturned(id int) (err error)
	GetBorrowedBookData(memberID, bookID int) (res BorrowedBooks, err error)
}

type ServiceInterface interface {
	InsertListOfBooks(input []Books) (err error)
	BorrowBooks(memberCode, bookCode string) (book Books, isServerErr bool, err error)
	ReturnBook(memberCode, bookCode string) (isServerErr bool, err error)
}
