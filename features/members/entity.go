package members

type Member struct {
	ID   int
	Code string
	Name string
}

type ListOfMembers struct {
	Member         Member
	BorrowedAmount int
}

type RepositoryInterface interface {
	ListMembersWithBorrowedAmount() (res []ListOfMembers, err error)
	InsertListOfMembers(input []Member) (err error)
}

type ServiceInterface interface {
	ListMembersWithBorrowedAmount() (res []ListOfMembers, err error)
	InsertListOfMembers(input []Member) (err error)
}
