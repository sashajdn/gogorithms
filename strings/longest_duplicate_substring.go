package strings

// LongestDuplicateSubstring_HashMap ...
//
// T -> O(s ** s)
// S -> O(s ** s)
func LongestDuplicateSubstring_HashMap(s string) int {
	var (
		hm           = map[string]struct{}{}
		longestSoFar int
	)

	for j := 0; j < len(s); j++ {
		for i := j; i < len(s); i++ {
			key := string(s[j : i+1])
			if _, ok := hm[key]; ok {
				longestSoFar = max(longestSoFar, len(key))
				continue
			}

			hm[key] = struct{}{}
		}
	}

	return longestSoFar
}

// LongestDuplicateSubstring_TrieBinarySearch ...
//
// T -> O(nk * log(n))
// S -> O(n)
func LongestDuplicateSubstring_TrieBinarySearch(s string) int {
	var pt = &PrefixTrie{
		directory:   make(map[rune]*PrefixTrie),
		isEndOfWord: false,
	}

	// T -> O(log(n) * nk)
	// S -> O(n)
	var left, right = 0, len(s) - 1
	for left < right {
		mid := (left + right) / 2

		if isSlidingWindowInTrie(pt, s, mid) {
			left = mid + 1
			continue
		}

		right = mid - 1
	}

	return left
}

// PrefixTrie ...
type PrefixTrie struct {
	directory   map[rune]*PrefixTrie
	isEndOfWord bool
}

// Insert ...
//
// T -> O(n)
// S -> O(n)
func (p *PrefixTrie) Insert(s string) bool {
	var (
		current = p
		found   bool
	)

	for i, c := range s {
		v, ok := current.directory[c]
		switch {
		case ok:
			found = (v.isEndOfWord && i == len(s)-1)

			if i == len(s)-1 {
				v.isEndOfWord = true
			}

		default:
			current.directory[c] = &PrefixTrie{
				directory:   make(map[rune]*PrefixTrie),
				isEndOfWord: i == len(s)-1,
			}
		}

		current = current.directory[c]
	}

	return found
}

// isSlidingWindowInTrie ...
//
// T -> O(nk)
// S -> O(n)
func isSlidingWindowInTrie(pt *PrefixTrie, s string, windowSize int) bool {
	// T -> O(nk)
	// S -> O(n)
	var (
		left, right = 0, windowSize - 1
		found       bool
	)
	for right <= len(s)-1 {
		substring := s[left : right+1]

		// T -> O(k)
		found = found || pt.Insert(substring)

		left++
		right++
	}

	return found
}
