package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	members "eigen-backend-test-case/features/members"
	middleware "eigen-backend-test-case/middleware"
	responses "eigen-backend-test-case/utils/responses"

	"github.com/julienschmidt/httprouter"
)

type membersDelivery struct {
	router  *httprouter.Router
	service members.ServiceInterface
}

func NewMembersDelivery(router *httprouter.Router, service members.ServiceInterface) {
	handler := &membersDelivery{
		router:  router,
		service: service,
	}

	router.GET("/api/v1/members/list_of_members", middleware.Cors(handler.ListMembersWithBorrowedAmount))
	router.POST("/api/v1/members/insert_list_of_members", middleware.Cors(handler.InsertListOfMembers))
}

func (d *membersDelivery) ListMembersWithBorrowedAmount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: list members with borrowed amount")

	members, err := d.service.ListMembersWithBorrowedAmount()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	response := toListOfMembersResponse(members)

	log.Printf(">>> response: list members with borrowed amount, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessWithDataResponse(response, "list of members with borrowed amount"))
}

func (d *membersDelivery) InsertListOfMembers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("<<< receive: insert list of members")

	var request []InsertMembersRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error(), r.RemoteAddr)
		return
	}

	toMembers := toInsertMembersRequest(request)

	err = d.service.InsertListOfMembers(toMembers)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), r.RemoteAddr)
		return
	}

	log.Printf(">>> response: insert list of members, %d - %s\n", http.StatusOK, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.SuccessResponse("members stored in database"))
}
