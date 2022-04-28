package strings

import (
	"fmt"
	"strings"
)

// GroupAnagrams ...
//
// T -> O(n * w)
// S -> O(n * w)
func GroupAnagrams(words []string) [][]string {
	var hm = make(map[string][]string, len(words))
	for _, word := range words {
		var counter = make(Counter, len(word))
		for _, r := range word {
			counter[r]++
		}

		key := counter.Hash()
		hm[key] = append(hm[key], word)
	}

	var output = make([][]string, 0, 0)
	for _, group := range hm {
		output = append(output, group)
	}

	return output
}

type Counter map[rune]int

func (c Counter) Hash() string {
	var sb strings.Builder
	for r := 'a'; r <= 'z'; r++ {
		if count, ok := c[r]; ok {
			sb.WriteString(fmt.Sprintf("#%v%d", r, count))
		}
	}

	return sb.String()
}
