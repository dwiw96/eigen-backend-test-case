package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	books "eigen-backend-test-case/features/books"
	middleware "eigen-backend-test-case/middleware"
	responses "eigen-backend-test-case/utils/responses"

	"github.com/julienschmidt/httprouter"
)

type booksDelivery struct {
	router  *httprouter.Router
	service books.ServiceInterface
}

func NewbooksDelivery(router *httprouter.Router, service books.ServiceInterface) {
	handler := &booksDelivery{
		router:  router,
		service: service,
	}

	router.POST("/api/v1/books/insert_list_of_books", middleware.Cors(handler.InsertListOfBooks))
	router.POST("/api/v1/books/borrow_book", middleware.Cors(handler.BorrowBook))
	router.PUT("/api/v1/books/return_book", middleware.Cors(handler.ReturnBook))
	router.GET("/api/v1/books/list_of_existing_books", middleware.Cors(handler.ListExistingBook))
}

func (d *booksDelivery) InsertListOfBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: insert list of books")

	var request []BooksRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error(), r.RemoteAddr)
		return
	}

	toBooks := toBooksDelivery(request)

	err = d.service.InsertListOfBooks(toBooks)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	log.Printf(">>> response: insert list of books, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessResponse("books stored in database"))
}

func (d *booksDelivery) BorrowBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: borrow book")

	var request BorrowBookRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error(), r.RemoteAddr)
		return
	}

	bookData, isServerError, err := d.service.BorrowBooks(request.MemberCode, request.BookCode)
	if err != nil {
		if !isServerError {
			responses.ErrorJSON(w, http.StatusBadRequest, "Bad Request", err.Error(), r.RemoteAddr)
			return
		}
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	response := toBookResponse(bookData)

	log.Printf(">>> response: borrow book success, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessWithDataResponse(response, "borrowed book data"))
}

func (d *booksDelivery) ReturnBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: return book")

	var request BorrowBookRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error(), r.RemoteAddr)
		return
	}

	isServerErr, err := d.service.ReturnBook(request.MemberCode, request.BookCode)
	if err != nil {
		if !isServerErr {
			responses.ErrorJSON(w, http.StatusBadRequest, "Bad Request", err.Error(), r.RemoteAddr)
			return
		}
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	log.Printf(">>> response: return book, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessResponse("book is returned"))
}

func (d *booksDelivery) ListExistingBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: list existing book")

	existingBooks, err := d.service.ListExistingBooks()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	response := toListExistingBooks(existingBooks)

	log.Printf(">>> response: list existing books, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessWithDataResponse(response, "list of existing books"))
}
