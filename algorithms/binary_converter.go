package algorithms

import "math"

const (
	zero = '0'
	one  = '1'
)

// BinaryConverter : O(T) -> O(n), O(S) -> O(n)
func BinaryConverter(s string) int {
	var total int
	for i, r := range reverse(s) {
		switch r {
		case one:
			total += int(math.Pow(2, float64(i)))
		}
	}
	return total
}

func reverse(s string) string {
	runes := []rune(s)
	return string(rev(runes))
}

func rev(x []rune) []rune {
	if len(x) == 0 {
		return []rune{}
	}
	return append(rev(x[1:]), x[0])
}
