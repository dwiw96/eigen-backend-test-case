package repository

import (
	"context"
	"os"
	"testing"
	"time"

	config "eigen-backend-test-case/config"
	members "eigen-backend-test-case/features/members"
	postgres "eigen-backend-test-case/utils/driver/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var repoTest members.RepositoryInterface

func TestMain(m *testing.M) {
	os.Setenv("DB_USERNAME", "dwiwahyudi")
	os.Setenv("DB_PASSWORD", "eigen123")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "eigen_book_project")

	envConfig := &config.EnvConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	db := postgres.InitDB(envConfig)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	repoTest = NewMembersRepository(db, ctx)

	os.Exit(m.Run())
}

func TestListMembersWithBorrowedAmount(t *testing.T) {
	tests := []struct {
		name string
		ans  []members.ListOfMembers
		err  bool
	}{
		{
			name: "success1",
			ans: []members.ListOfMembers{
				{
					Member: members.Member{
						ID:   1,
						Code: "M001",
						Name: "Angga",
					},
					BorrowedAmount: 0,
				}, {
					Member: members.Member{
						ID:   2,
						Code: "M002",
						Name: "Ferry",
					},
					BorrowedAmount: 0,
				}, {
					Member: members.Member{
						ID:   3,
						Code: "M003",
						Name: "Putri",
					},
					BorrowedAmount: 0,
				}, {
					Member: members.Member{
						ID:   4,
						Code: "M004",
						Name: "Dwi",
					},
					BorrowedAmount: 0,
				}, {
					Member: members.Member{
						ID:   5,
						Code: "M005",
						Name: "Wahyu",
					},
					BorrowedAmount: 1,
				}, {
					Member: members.Member{
						ID:   6,
						Code: "M006",
						Name: "Yudi",
					},
					BorrowedAmount: 0,
				},
			},
			err: false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.ListMembersWithBorrowedAmount()
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				require.Error(t, err)
			}
		})
	}
}
