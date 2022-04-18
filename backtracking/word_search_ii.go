package backtracking

// WordSearchII ...
//
// T -> O(N * 3^L) where `N` is the number of cells in the board & L the max length of words.
// S -> O(w) where `w` is the number of chars in the prefix trie.
func WordSearchII(board [][]byte, words []string) []string {
	var pt = &PrefixTrie{
		h:         make(map[rune]*PrefixTrie),
		EndOfWord: false,
	}
	for _, word := range words {
		pt.AddWord(word)
	}

	var (
		foundWords = make([]string, 0, len(words))
		set        = map[string]struct{}{}
	)

	var backtrack func(j, i int, currentPT *PrefixTrie, currentWord string)
	backtrack = func(j, i int, currentPT *PrefixTrie, currentWord string) {
		if rune(board[j][i]) == '#' {
			return
		}

		current := board[j][i]
		board[j][i] = byte('#')
		defer func() {
			board[j][i] = current
		}()

		currentChar := rune(current)
		currentPT = currentPT.SearchChar(currentChar)
		if currentPT == nil {
			return
		}

		currentWord += string(currentChar)
		if currentPT.EndOfWord {
			_, ok := set[currentWord]
			switch {
			case ok:
			default:
				foundWords = append(foundWords, currentWord)
				set[currentWord] = struct{}{}
			}
		}

		for _, neighbour := range fetchNeighbours(j, i, board) {
			dj, di := neighbour[0], neighbour[1]
			backtrack(dj, di, currentPT, currentWord)
		}
	}

	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			backtrack(j, i, pt, "")
		}
	}

	return foundWords
}

func fetchNeighbours(j, i int, board [][]byte) [][]int {
	var (
		neighbours = make([][]int, 0, 3)
		directions = [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}
	)
	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[1]

		if dj < 0 || dj > len(board)-1 {
			continue
		}
		if di < 0 || di > len(board[0])-1 {
			continue
		}

		if rune(board[dj][di]) == '#' {
			continue
		}

		neighbours = append(neighbours, []int{dj, di})
	}
	return neighbours
}

type PrefixTrie struct {
	h         map[rune]*PrefixTrie
	EndOfWord bool
}

func (p *PrefixTrie) AddWord(word string) {
	var current = p
	for i, r := range word {
		if v, ok := current.h[r]; ok {
			current = v
			if current.EndOfWord {
				continue
			}

			current.EndOfWord = i == len(word)-1
			continue
		}

		latest := &PrefixTrie{
			h:         make(map[rune]*PrefixTrie),
			EndOfWord: i == len(word)-1,
		}

		current.h[r] = latest
		current = current.h[r]
	}
}

func (p *PrefixTrie) SearchChar(char rune) *PrefixTrie {
	if pt, ok := p.h[char]; ok {
		return pt
	}

	return nil
}
