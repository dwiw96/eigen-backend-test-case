package repository

import (
	"context"
	"os"
	"testing"
	"time"

	config "eigen-backend-test-case/config"
	books "eigen-backend-test-case/features/books"
	postgres "eigen-backend-test-case/utils/driver/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var repoTest books.RepositoryInterface

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

	repoTest = NewBooksRepository(db, ctx)

	os.Exit(m.Run())
}

func TestInsertListOfBooks(t *testing.T) {
	tests := []struct {
		name  string
		input []books.Books
		err   bool
	}{
		{
			name: "success1",
			input: []books.Books{
				{
					Code:   "JK-45",
					Title:  "Harry Potter",
					Author: "J.K Rowling",
					Stock:  1,
				}, {
					Code:   "SHR-1",
					Title:  "A Study in Scarlet",
					Author: "Arthur Conan Doyle",
					Stock:  1,
				}, {
					Code:   "TW-11",
					Title:  "Twilight",
					Author: "Stephenie Meyer",
					Stock:  1,
				}, {
					Code:   "HOB-83",
					Title:  "The Hobbit, or There and Back Again",
					Author: "J.R.R. Tolkien",
					Stock:  1,
				}, {
					Code:   "NRN-7",
					Title:  "The Lion, the Witch and the Wardrobe",
					Author: "C.S. Lewis",
					Stock:  1,
				},
			},
			err: false,
		},
		{
			name: "error1",
			input: []books.Books{
				{
					Code:   "JK-45",
					Title:  "",
					Author: "J.K Rowling",
					Stock:  1,
				}, {
					Code:  "SHR-1",
					Title: "A Study in Scarlet",
					Stock: 1,
				}, {
					Code:   "TW-11",
					Title:  "Twilight",
					Author: "Stephenie Meyer",
					Stock:  0,
				},
			},
			err: true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := repoTest.InsertListOfBooks(v.input)
			if !v.err {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestGetBookData(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ans   books.Books
		err   bool
	}{
		{
			name:  "success1",
			input: "JK-45",
			ans: books.Books{
				ID:     1,
				Code:   "JK-45",
				Title:  "Harry Potter",
				Author: "J.K Rowling",
				Stock:  1,
			},
			err: false,
		}, {
			name:  "success2",
			input: "TW-11",
			ans: books.Books{
				ID:     3,
				Code:   "TW-11",
				Title:  "Twilight",
				Author: "Stephenie Meyer",
				Stock:  1,
			},
			err: false,
		}, {
			name:  "error1",
			input: "ABX-56",
			err:   true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.GetBookData(v.input)
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				require.Error(t, err)
				assert.Empty(t, res)
			}
		})
	}
}

func TestCheckIfBookIsAvailable(t *testing.T) {
	tests := []struct {
		name  string
		input int
		ans   int
		err   bool
	}{
		{
			name:  "success1",
			input: 1,
			ans:   0,
			err:   false,
		}, {
			name:  "success2",
			input: 2,
			ans:   0,
			err:   false,
		}, {
			name:  "success3",
			input: 3,
			ans:   0,
			err:   false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.CheckIfBookIsAvailable(v.input)
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				require.Error(t, err)
				assert.Zero(t, res)
			}
		})
	}
}

func TestGetMemberData(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ans   books.Member
		err   bool
	}{
		{
			name:  "success1",
			input: "M001",
			ans: books.Member{
				ID:   1,
				Code: "M001",
				Name: "Angga",
			},
			err: false,
		}, {
			name:  "success2",
			input: "M002",
			ans: books.Member{
				ID:   2,
				Code: "M002",
				Name: "Ferry",
			},
			err: false,
		}, {
			name:  "error1",
			input: "ABX-56",
			err:   true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.GetMemberData(v.input)
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				assert.Empty(t, res)
			}
		})
	}
}

func TestCheckMemberBorrowedBooks(t *testing.T) {
	tests := []struct {
		name  string
		input int
		ans   int
		err   bool
	}{
		{
			name:  "success1",
			input: 1,
			ans:   0,
			err:   false,
		}, {
			name:  "success2",
			input: 2,
			ans:   0,
			err:   false,
		}, {
			name:  "success3",
			input: 3,
			ans:   0,
			err:   false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.CheckMemberBorrowedBooks(v.input)
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				require.Error(t, err)
				assert.Zero(t, res)
			}
		})
	}
}

func TestCheckIfMemberPenalized(t *testing.T) {
	tests := []struct {
		name  string
		input int
		ans   bool
		err   bool
	}{
		{
			name:  "success1",
			input: 1,
			ans:   false,
			err:   false,
		}, {
			name:  "success2",
			input: 2,
			ans:   false,
			err:   false,
		}, {
			name:  "success3",
			input: 3,
			ans:   false,
			err:   false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := repoTest.CheckIfMemberPenalized(v.input)
			if !v.err {
				require.NoError(t, err)
				assert.Equal(t, v.ans, res)
			} else {
				require.Error(t, err)
				assert.Zero(t, res)
			}
		})
	}
}

func TestInsertBorrowedBook(t *testing.T) {
	tests := []struct {
		name     string
		bookID   int
		memberID int
		err      bool
	}{
		{
			name:     "success1",
			bookID:   6,
			memberID: 4,
			err:      false,
		}, {
			name:     "success2",
			bookID:   7,
			memberID: 5,
			err:      false,
		}, {
			name:     "success3",
			bookID:   7,
			memberID: 4,
			err:      false,
		}, {
			name:     "error1",
			bookID:   100,
			memberID: 5,
			err:      true,
		}, {
			name:     "success3",
			bookID:   8,
			memberID: 100,
			err:      true,
		}, {
			name:     "success3",
			bookID:   800,
			memberID: 100,
			err:      true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := repoTest.InsertBorrowedBook(v.bookID, v.memberID)
			if !v.err {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
