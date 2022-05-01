package dynamic

type Transaction struct {
	ID    string
	Size  int
	Value int
}

// KnapsackTransactions ...
//
// T -> O(t * b)
// S -> O(t * b)
func KnapsackTransactions(transactions []*Transaction, blockSize int) []*Transaction {
	var dp = make([][]int, 0, len(transactions)+1)
	for i := 0; i < len(transactions)+1; i++ {
		dp = append(dp, make([]int, blockSize+1))
	}

	for j := 1; j < len(transactions)+1; j++ {
		transaction := transactions[j-1]
		for i := 1; i < blockSize+1; i++ {
			switch {
			case i < transaction.Size:
				dp[j][i] = dp[j-1][i]
			default:
				dp[j][i] = max(dp[j-1][i], dp[j-1][i-transaction.Size]+transaction.Value)
			}
		}
	}

	var (
		includedTransactions []*Transaction
		capacityLeft         = blockSize
	)
	for j := len(transactions); j >= 0; j-- {
		if capacityLeft <= 0 {
			break
		}

		if dp[j][capacityLeft] == dp[j-1][capacityLeft] {
			continue
		}

		includedTransactions = append(includedTransactions, transactions[j-1])
		capacityLeft -= transactions[j-1].Size
	}

	var left, right int
	for left < right {
		includedTransactions[left], includedTransactions[right] = includedTransactions[right], includedTransactions[left]
		left++
		right--
	}

	return includedTransactions
}
