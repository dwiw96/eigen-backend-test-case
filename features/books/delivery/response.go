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
