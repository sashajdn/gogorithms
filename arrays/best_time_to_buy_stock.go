package arrays

// BestTimeToBuyStock ...
// T -> O(n) where `n` is the number of prices.
// S -> O(1)
func BestTimeToBuyStock(prices []int) int {
	var (
		maxProfitSoFar int
		minSoFar       = prices[0]
	)

	for _, price := range prices[1:] {
		maxProfitSoFar = max(maxProfitSoFar, price-minSoFar)
		minSoFar = min(minSoFar, price)
	}

	return maxProfitSoFar
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
