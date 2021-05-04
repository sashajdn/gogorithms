package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTournamentWinner(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		inputCompetitions [][]string
		inputResults      []int
		expectedResult    string
	}{
		{
			name: "3_teams",
			inputCompetitions: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			inputResults:   []int{0, 0, 1},
			expectedResult: "C",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := TournamentWinner(tc.inputCompetitions, tc.inputResults)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}
