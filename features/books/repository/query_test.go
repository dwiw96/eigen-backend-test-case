package repository

import (
	"context"
	"os"
	"testing"
	"time"

	config "eigen-backend-test-case/config"
	books "eigen-backend-test-case/features/books"
	postgres "eigen-backend-test-case/utils/driver/postgres"

	// "github.com/stretchr/testify/assert"
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
