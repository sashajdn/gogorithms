package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var editDistanceChecks = []func(string, string) int{
	EditDistance,
	EditDistanceII_TopDown2D,
	EditDistanceII_TopDown1D,
}

func TestEditDistance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		a, b                 string
		expectedEditDistance int
	}{
		{
			name:                 "example_one",
			a:                    "horse",
			b:                    "ros",
			expectedEditDistance: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range editDistanceChecks {
				res := c(tt.a, tt.b)

				assert.Equal(t, tt.expectedEditDistance, res)
			}
		})
	}
}
