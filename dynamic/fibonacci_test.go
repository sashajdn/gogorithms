package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibonnaciChecks = []func(n int) int{
	Fibonacci_BottomUp,
	Fibonacci_TopDown,
	Fibonacci_TopDownPointers,
}

func TestFibonacci(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		n              int
		expectedOutput int
	}{
		{
			name:           "0",
			n:              0,
			expectedOutput: 0,
		},
		{
			name:           "1",
			n:              1,
			expectedOutput: 1,
		},
		{
			name:           "2",
			n:              2,
			expectedOutput: 1,
		},
		{
			name:           "14",
			n:              14,
			expectedOutput: 377,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range fibonnaciChecks {
				assert.Equal(t, tt.expectedOutput, c(tt.n))
			}
		})
	}
}
