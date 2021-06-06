package algorithms

import "fmt"

// RunLengthEncoding : O(T) -> O(n), O(S) -> O(n)
func RunLengthEncoding(str string) string {
	var (
		res     string
		curr    rune
		howMany int
	)
	for _, r := range str {
		if howMany == 0 {
			curr = r
			howMany++
			continue
		}
		if r != curr {
			res += fmt.Sprintf("%v%s", howMany, string(curr))
			curr = r
			howMany = 1
			continue
		}
		howMany++
		if howMany > 9 {
			res += fmt.Sprintf("%v%s", 9, string(r))
			howMany = 1
		}
	}
	res += fmt.Sprintf("%v%s", howMany, string(curr))
	return res
}
