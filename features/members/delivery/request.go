package delivery

import members "eigen-backend-test-case/features/members"

type InsertMembersRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func toInsertMembersRequest(input []InsertMembersRequest) (res []members.Member) {
	for _, v := range input {
		var temp members.Member
		temp.Code = v.Code
		temp.Name = v.Name

		res = append(res, temp)
	}
	return
}
