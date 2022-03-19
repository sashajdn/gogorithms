package trie

// PrefixTrie ...
type PrefixTrie struct {
	children   map[rune]*PrefixTrie
	IsComplete bool
}

// Insert ...
// T -> O(n) where n is the length of the word in the worst case.
// S -> O(n) in the worst case.
func (t *PrefixTrie) Insert(word string) {
	var current = t
	for _, r := range word {
		current = current.addChild(r, r == rune(word[len(word)-1]))
	}
}

// Search ...
//
// T -> O(n) where n is the length of the word in the worst case.
// S -> O(1)
func (t *PrefixTrie) Search(word string) bool {
	var current = t
	for _, r := range word {
		var ok bool
		current, ok = current.children[r]
		if !ok {
			return false
		}
	}
	if current == nil {
		return false
	}

	return current.IsComplete
}

// StartsWith ...
//
// T -> O(n) where n is the length of the word in the worst case.
// S -> O(1)
func (t *PrefixTrie) StartsWith(prefix string) bool {
	var current = t
	for _, r := range prefix {
		var ok bool
		current, ok = current.children[r]
		if !ok {
			return false
		}
	}
	return true
}

func (t *PrefixTrie) addChild(char rune, endOfWord bool) *PrefixTrie {
	child, ok := t.children[char]
	if ok {
		child.IsComplete = endOfWord
		return child
	}

	newChild := NewPrefixTrie(endOfWord)
	t.children[char] = newChild
	return newChild
}

func NewPrefixTrie(isComplete bool) *PrefixTrie {
	return &PrefixTrie{
		children:   make(map[rune]*PrefixTrie),
		IsComplete: isComplete,
	}
}
