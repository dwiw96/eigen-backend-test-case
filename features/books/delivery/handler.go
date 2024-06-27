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
