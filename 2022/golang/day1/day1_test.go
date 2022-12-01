package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day1(t *testing.T) {
	type input struct {
		items    []string
		topCount int
	}
	tests := map[string]struct {
		input    input
		expected int
	}{
		"test1": {
			input: input{
				readInput("input_test1.txt"),
				3,
			},
			expected: 45000,
		},
		"final test": {
			input: input{
				readInput("input.txt"),
				3,
			},
			expected: 210367,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, maxCalories(tt.input.items, tt.input.topCount))
		})
	}
}
