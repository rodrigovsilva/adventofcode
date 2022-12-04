package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_totalScore(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected int
	}{
		"test 1": {
			input:    readInput("input_test1.txt"),
			expected: 12,
		},
		"test final": {
			input:    readInput("input.txt"),
			expected: 13509,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, totalScore(tt.input))
		})

	}
}
