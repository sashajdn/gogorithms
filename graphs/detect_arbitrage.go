package graphs

import "math"

// DetectArbitrage_Improved ...
//
// T -> O(n ** 3)
// S -> O(n)
func DetectArbitrage_Improved(exchangeRates [][]float64) bool {
	var profits = make([]float64, 0, len(exchangeRates))
	for i := 0; i < len(profits); i++ {
		profits[i] = math.MaxInt
	}
	profits[0] = 0

	// Relaxation.
	for n := 0; n < len(exchangeRates); n++ {
		for j := 0; j < len(exchangeRates); j++ {
			for i := 0; i < len(exchangeRates[0]); i++ {
				from, to := j, i
				profit := negativeLog(exchangeRates[j][i])

				if profits[from] == math.MaxInt {
					profits[to] = profits[from] + profit
					continue
				}

				if profits[from]+profit < profits[to] {
					profits[to] = profits[from] + profit
				}
			}
		}
	}

	// Propagate negative cycles.
	for n := 0; n < len(exchangeRates); n++ {
		for j := 0; j < len(exchangeRates); j++ {
			for i := 0; i < len(exchangeRates); i++ {
				from, to := j, i
				if profits[from] == math.MinInt {
					profits[to] = math.MinInt
					continue
				}

				profit := negativeLog(exchangeRates[j][i])
				if profits[from]+profit < profits[to] {
					profits[to] = math.MinInt
					continue
				}
			}
		}
	}

	// Find negative cycle.
	for _, p := range profits {
		if p == math.MinInt {
			return true
		}
	}

	return false
}

//  ArbitrageEdge ...
type ArbitrageEdge struct {
	To, From int
	Cost     float64
}

// DetectArbitrage ...
//
// T -> O(v * v * v) where `v` is the number of exchanges.
// S -> O(v * v)
func DetectArbitrage(exchangeRates [][]float64) bool {
	var edges []*ArbitrageEdge
	for j := 0; j < len(exchangeRates); j++ {
		for i := 0; i < len(exchangeRates[0]); i++ {
			if i == j {
				continue
			}

			edges = append(edges, &ArbitrageEdge{
				From: j,
				To:   i,
				Cost: negativeLog(exchangeRates[j][i]),
			})
		}
	}

	numberOfExchanges := len(exchangeRates)

	distances := make([]float64, numberOfExchanges)
	for i := 0; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}
	distances[0] = 0 // We can start at any currency since we have a strongly connected graph.

	for i := 0; i < numberOfExchanges; i++ {
		for _, edge := range edges {
			if distances[edge.From] == math.MaxInt {
				distances[edge.To] = edge.Cost
				continue
			}

			distances[edge.To] = floatMin(distances[edge.To], distances[edge.From]+edge.Cost)

		}
	}

	for i := 0; i < numberOfExchanges; i++ {
		for _, edge := range edges {
			if distances[edge.From]+edge.Cost < distances[edge.To] {
				return true
			}
		}
	}

	return false
}

func negativeLog(value float64) float64 {
	return -1 * math.Log(value)
}

func floatMin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
