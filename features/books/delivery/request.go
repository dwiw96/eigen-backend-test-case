package delivery

import (
	books "eigen-backend-test-case/features/books"
)

type BooksRequest struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"`
}

func toBooksDelivery(input []BooksRequest) (res []books.Books) {
	for _, v := range input {
		var temp books.Books
		temp.Code = v.Code
		temp.Title = v.Title
		temp.Author = v.Author
		temp.Stock = v.Stock

		res = append(res, temp)
	}

	return
}

type BorrowBookRequest struct {
	BookCode   string `json:"book_code"`
	MemberCode string `json:"member_code"`
}
