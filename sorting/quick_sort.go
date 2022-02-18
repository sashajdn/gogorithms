package sorting

// QuickSort ...
// T -> O(nlog(n)) Best Case, Worst Case O(n**2)
// S -> O(log(n))
func QuickSort(array []int) {
	quickSort(&array)
}

// quickSort ...
func quickSort(array *[]int) {
	if len(*array) < 2 {
		return
	}

	if len(*array) == 2 {
		if (*array)[0] > (*array)[1] {
			swap(array, 0, 1)
		}
		return
	}

	pivot := 0
	l, r := pivot+1, len(*array)-1

	for r >= l {
		if (*array)[l] > (*array)[pivot] && (*array)[r] < (*array)[pivot] {
			swap(array, l, r)
		}

		if (*array)[l] <= (*array)[pivot] {
			l++
		}

		if (*array)[r] >= (*array)[pivot] {
			r--
		}

	}
	swap(array, pivot, r)

	lsa, rsa := (*array)[:r], (*array)[r+1:]

	quickSort(&lsa)
	quickSort(&rsa)
}

func swap(array *[]int, i, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
