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
