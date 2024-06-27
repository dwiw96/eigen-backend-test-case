package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	initConfig()
}

func TestGetConfig(t *testing.T) {
	ans := &EnvConfig{
		SERVER_PORT: ":8080",
		DB_USERNAME: "dwiwahyudi",
		DB_PASSWORD: "eigen123",
		DB_HOST:     "localhost",
		DB_PORT:     "5432",
		DB_NAME:     "eigen_book_project",
		// ABSOLUTE_PATH: "/home/dwiw22/107/inventory-system",
	}
	res := GetConfig()
	t.Log("RES =", res)
	require.NotNil(t, res)
	assert.Equal(t, ans, res)
}
