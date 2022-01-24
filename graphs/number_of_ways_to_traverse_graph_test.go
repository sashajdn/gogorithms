package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfWaysToTraverseAGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		width, height  int
		expectedOutput int
	}{
		{},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := NumberOfWaysToTraverseGraph(tt.width, tt.height)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}

}
