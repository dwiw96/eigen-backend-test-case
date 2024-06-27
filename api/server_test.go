package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetUpRouter(t *testing.T) {
	router := SetUpRouter()

	require.NotNil(t, router)
}

func TestStartServer(t *testing.T) {
	router := SetUpRouter()
	require.NotNil(t, router)

	StartServer(":8080", router)
}
