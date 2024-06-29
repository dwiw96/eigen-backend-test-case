package delivery

import (
	books "eigen-backend-test-case/features/books"
)

type BookResponse struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func toBookResponse(input books.Books) (res BookResponse) {
	res.Code = input.Code
	res.Title = input.Title
	res.Author = input.Author

	return
}

type ListExistingBookResponse struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"`
}

func toListExistingBooks(input []books.Books) (res []ListExistingBookResponse) {
	for _, v := range input {
		var temp ListExistingBookResponse
		temp.Code = v.Code
		temp.Title = v.Title
		temp.Author = v.Author
		temp.Stock = v.Stock

		res = append(res, temp)
	}

	return
}
