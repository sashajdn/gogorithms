package graphs

import (
	"math"
)

// ExchangeRate ...
type ExchangeRate struct {
	Source  string
	Target  string
	BestBid float64
	BestAsk float64
}

// ExchangeEdge ...
type ExchangeEdge struct {
	To, From int
	Cost     float64
}

// ExchangeResult ...
type ExchangeResult struct {
	To, From string
	Rate     float64
	Path     []string
}

// BestExchangeRatePath ...
//
// O -> (n ** 3) where `n` is the number of possible exchange rates.
// O -> (n ** 2)
func BestExchangeRatePath(sourceCurrency, targetCurrency string, exchangeRates []*ExchangeRate) *ExchangeResult {
	var (
		edges          []*ExchangeEdge
		currencyToNode = map[string]int{}
		nodeToCurrency = map[int]string{}
	)

	// Build graph of rates as edges.
	// T -> O(n) worst case when we have a complete graph: here `n` is the number of exchanges.
	// S -> O(n ** 2)
	for _, rate := range exchangeRates {
		if _, ok := currencyToNode[rate.Source]; !ok {
			nodeID := len(currencyToNode)

			currencyToNode[rate.Source] = nodeID
			nodeToCurrency[nodeID] = rate.Source

		}
		if _, ok := currencyToNode[rate.Target]; !ok {
			nodeID := len(currencyToNode)

			currencyToNode[rate.Target] = nodeID
			nodeToCurrency[nodeID] = rate.Target
		}

		edges = append(edges, &ExchangeEdge{
			Cost: normalizer(rate.BestAsk / rate.BestBid),
			From: currencyToNode[rate.Source],
			To:   currencyToNode[rate.Target],
		})
	}

	// Bellman-Ford.
	// T -> O(n ** 3)
	var (
		distances = make([]float64, len(currencyToNode))
		previous  = make([]int, len(currencyToNode))
	)
	for i := 0; i < len(currencyToNode); i++ {
		distances[i] = math.MaxInt
	}
	distances[currencyToNode[sourceCurrency]] = 0

	// Relax nodes.
	for i := 0; i < len(currencyToNode); i++ {
		for _, edge := range edges {
			if distances[edge.From] == math.MaxInt {
				distances[edge.To] = edge.Cost
				previous[edge.To] = edge.From
				continue
			}

			distances[edge.From] = floatMin(distances[edge.To], distances[edge.From]+edge.Cost)
		}
	}

	// Search for negative cycles (we shouldn't have this if we don't have arbitrages).
	for i := 0; i < len(currencyToNode); i++ {
		for _, edge := range edges {
			if distances[edge.From]+edge.Cost < distances[edge.To] {
				distances[edge.To] = math.MinInt
				previous[edge.To] = -1
			}
		}
	}
	minCostToTarget := distances[currencyToNode[targetCurrency]]

	// Here we have a negative cycle.
	if minCostToTarget == math.MinInt {
		return nil
	}

	// Rebuild path taken for best exchange.
	var (
		path       = []string{targetCurrency}
		currentIdx = currencyToNode[targetCurrency]
	)
	for i := 0; i < len(previous); i++ {
		previousIdx := previous[currentIdx]
		if previousIdx < 0 {
			path = []string{}
			break
		}

		path = append(path, nodeToCurrency[previousIdx])
		currentIdx = previousIdx
	}
	reverse(path)

	return &ExchangeResult{
		From: sourceCurrency,
		To:   targetCurrency,
		Rate: math.Exp(abs(minCostToTarget)), // inverse our normalization function.
		Path: path,
	}
}

func normalizer(value float64) float64 {
	return -1 * math.Log(value)
}

func abs(value float64) float64 {
	if value < 0 {
		return -1 * value
	}

	return value
}

func reverse(array []string) {
	var (
		i = 0
		j = len(array) - 1
	)
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
}
