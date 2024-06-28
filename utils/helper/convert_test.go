package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatGoTime(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		err   bool
	}{
		{
			name:  "success1",
			input: time.Now(),
			err:   false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := FormatGoTime(v.input)
			assert.NotEmpty(t, res)
			t.Log(res)
		})
	}
}

func TestParsingPgTime(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   bool
	}{
		{
			name:  "success1",
			input: "2024-06-28 09:45:55",
			err:   false,
		}, {
			name:  "success2",
			input: "2024-06-28 09:45:55",
			err:   false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := ParsingPgTime(v.input)
			if !v.err {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}

			t.Log(res)
		})
	}
}
