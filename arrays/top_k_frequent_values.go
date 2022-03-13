package arrays

type Frequency struct {
	Value int
	F     int
}

// T -> Best case O(n)
// S -> O(n)
func TopKFrequentValues(nums []int, k int) []int {
	var (
		frequencies    = map[int]int{}
		frequenciesArr = make([]*Frequency, 0, len(nums))
		output         = make([]int, 0, len(nums))
	)

	for _, num := range nums {
		frequencies[num]++
	}

	for k, v := range frequencies {
		frequenciesArr = append(frequenciesArr, &Frequency{
			Value: k,
			F:     v,
		})
	}

	topK := quickSelectTopK(frequenciesArr, 0, len(frequenciesArr)-1, len(frequenciesArr)-k)
	for _, f := range topK {
		output = append(output, f.Value)
	}

	return output
}

func quickSelectTopK(array []*Frequency, start, end, target int) []*Frequency {
	if start > end {
		panic("")
	}

	p := start
	l, r := start+1, end
	for l <= r {
		if array[l].F > array[p].F && array[r].F < array[p].F {
			swapTopK(array, l, r)
		}

		if array[l].F <= array[p].F {
			l++
		}

		if array[r].F >= array[p].F {
			r--
		}
	}
	swapTopK(array, r, p)

	if r == target {
		return array[target:]
	}

	if r < target {
		return quickSelectTopK(array, r+1, end, target)
	}

	return quickSelectTopK(array, start, r-1, target)
}

func swapTopK(array []*Frequency, i, j int) {
	array[i], array[j] = array[j], array[i]
}
