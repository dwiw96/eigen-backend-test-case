package delivery

import members "eigen-backend-test-case/features/members"

type ListOfMembersResponse struct {
	ID             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	BorrowedAmount int    `json:"borrowed_amount"`
}

func toListOfMembersResponse(input []members.ListOfMembers) (res []ListOfMembersResponse) {
	for _, v := range input {
		var temp ListOfMembersResponse
		temp.Code = v.Member.Code
		temp.Name = v.Member.Name
		temp.BorrowedAmount = v.BorrowedAmount

		res = append(res, temp)
	}

	return
}
