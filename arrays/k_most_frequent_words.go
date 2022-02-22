package arrays

func TopKFrequent(words []string, k int) []string {
	var frequencies = map[string]int{}
	for _, word := range words {
		frequencies[word]++
	}

	var frequencyTuples []*FrequencyPair
	for word, frequency := range frequencies {
		frequencyTuples = append(frequencyTuples, &FrequencyPair{
			Word:      word,
			Frequency: frequency,
		})
	}

	topKFrequencyPairs := topKFrequentWordsQuickSelect(frequencyTuples, 0, len(frequencyTuples)-1, len(frequencyTuples)-k)

	var topKFrequentWords = make([]string, 0, len(topKFrequencyPairs))
	for _, pair := range topKFrequencyPairs {
		topKFrequentWords = append(topKFrequentWords, pair.String())
	}

	return topKFrequentWords
}

func topKFrequentWordsQuickSelect(frequencyTuples []*FrequencyPair, start, end, k int) []*FrequencyPair {
	if k > len(frequencyTuples) {
		return frequencyTuples
	}

	pivot := start
	l, r := start+1, end
	for l <= r {
		pt, lt, rt := frequencyTuples[pivot], frequencyTuples[l], frequencyTuples[r]

		if lt.GreaterThan(pt) && rt.LessThan(pt) {
			frequencyPairSwap(frequencyTuples, l, r)
		}

		if !lt.GreaterThan(pt) {
			l++
		}

		if !rt.LessThan(pt) {
			r--
		}
	}
	frequencyPairSwap(frequencyTuples, pivot, r)

	if r == k {
		return frequencyTuples[k:]
	}

	if r < k {
		return topKFrequentWordsQuickSelect(frequencyTuples, r+1, end, k)
	}

	return topKFrequentWordsQuickSelect(frequencyTuples, start, r-1, k)
}

func frequencyPairSwap(array []*FrequencyPair, i, j int) {
	array[i], array[j] = array[j], array[i]
}

type FrequencyPair struct {
	Word      string
	Frequency int
}

func (f *FrequencyPair) LessThan(other *FrequencyPair) bool {
	if f.Frequency == other.Frequency {
		return f.Word < other.Word
	}

	return f.Frequency < other.Frequency
}

func (f *FrequencyPair) GreaterThan(other *FrequencyPair) bool {
	if f.Frequency == other.Frequency {
		return f.Word > other.Word
	}

	return f.Frequency > other.Frequency
}

func (f *FrequencyPair) String() string {
	return f.Word
}
