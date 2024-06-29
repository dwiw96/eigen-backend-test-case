package service

import (
	"context"
	members "eigen-backend-test-case/features/members"
)

type membersService struct {
	repo members.RepositoryInterface
	ctx  context.Context
}

func NewMembersService(repo members.RepositoryInterface, ctx context.Context) members.ServiceInterface {
	return &membersService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *membersService) ListMembersWithBorrowedAmount() (res []members.ListOfMembers, err error) {
	res, err = s.repo.ListMembersWithBorrowedAmount()
	if err != nil {
		return nil, err
	}

	return
}

func (s *membersService) InsertListOfMembers(input []members.Member) (err error) {
	err = s.repo.InsertListOfMembers(input)
	if err != nil {
		return err
	}

	return
}
