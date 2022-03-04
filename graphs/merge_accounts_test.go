package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mergeAccountFuncs = []func([][]string) [][]string{
	MergeAccounts,
	MergeAccounts_Improved,
}

func TestMergeAccounts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		accounts       [][]string
		expectedResult [][]string
	}{
		{
			name: "example_one",
			accounts: [][]string{
				{"John", "e1", "e2", "e3"},
				{"John", "e3", "e4"},
				{"Dave", "e6"},
			},
			expectedResult: [][]string{
				{"John", "e1", "e2", "e3", "e4"},
				{"Dave", "e6"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, f := range mergeAccountFuncs {
				res := f(tt.accounts)

				assert.Equal(t, tt.expectedResult, res)
			}
		})
	}
}
