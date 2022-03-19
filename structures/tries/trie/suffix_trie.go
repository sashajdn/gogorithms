package trie

// SuffixTrie ...
type SuffixTrie map[rune]*SuffixTrie

// NewSuffixTrie ...
//
// T -> O(1)
// S -> O(1)
func NewSuffixTrie() *SuffixTrie {
	return &SuffixTrie{}
}

// PopulateFromString ...
//
// T -> Worst case: O(n ** 2), where `n` is the length of the string.
// S -> O(n ** 2) where `n` is the length of the string.
func (t *SuffixTrie) PopulateFromString(s string) {
	for i := 0; i < len(s); i++ {
		var current = t
		for j := i; j < len(s); j++ {
			var r = rune(s[j])
			if _, ok := (*current)[r]; !ok {
				(*current)[r] = NewSuffixTrie()
			}

			(*current)[r] = NewSuffixTrie()
		}

		// Word termination signal.
		(*current)['*'] = nil
	}
}

// Contains ...
// T -> O(n) where `n` is the length of the suffix string to search for.
// S -> O(1)
func (t *SuffixTrie) Contains(s string) bool {
	var current = t
	for _, r := range s {
		var ok bool
		current, ok = (*current)[r]
		if !ok {
			return false
		}
	}

	if _, ok := (*current)['*']; ok {
		return true
	}

	return false
}
