package arrays

// T -> O(n), S -> O(n)
func arrayOfProductsBasic(array []int) []int {
	numZeros := 0
	zeroIndex := -1
	totalProduct := 1

	result := []int{}
	for i, val := range array {
		result = append(result, 0)
		if val == 0 {
			numZeros += 1
			zeroIndex = i
			continue
		}
		totalProduct *= val
	}

	if numZeros > 1 {
		return result
	}

	if numZeros == 1 {
		result[zeroIndex] = totalProduct
		return result
	}

	for i, val := range array {
		result[i] = totalProduct / val
	}
	return result
}
