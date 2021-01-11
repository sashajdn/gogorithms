package dynamic

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		a, b             string
		expectedDistance int
	}{
		{
			name:             "empty-strings",
			a:                "",
			b:                "",
			expectedDistance: 0,
		},
		{
			name:             "empty-string",
			a:                "abc",
			b:                "",
			expectedDistance: 3,
		},
		{
			name:             "strings-same-length",
			a:                "abc",
			b:                "deg",
			expectedDistance: 3,
		},
		{
			name:             "strings-diff-length",
			a:                "abct",
			b:                "abc",
			expectedDistance: 1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := LevenshteinDistance(tc.a, tc.b)
			assert(t, tc.expectedDistance, res)
		})
	}
}

func TestLongestString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		a, b            string
		expectedLongest string
	}{
		{
			name:            "same",
			a:               "a",
			b:               "b",
			expectedLongest: "a",
		},
		{
			name:            "a",
			a:               "ab",
			b:               "c",
			expectedLongest: "ab",
		},
		{
			name:            "b",
			a:               "a",
			b:               "bc",
			expectedLongest: "bc",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a, _ := longestString(tc.a, tc.b)
			assert(t, tc.expectedLongest, a)
		})
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	res := min(7, 3, 5)
	assert(t, 3, res)

	res = min(7, 3, 5, 0)
	assert(t, 0, res)
}

func assert(t *testing.T, expected interface{}, got interface{}) {
	if expected != got {
		t.Fatalf("expected -> %v, got -> %v", expected, got)
	}
}
