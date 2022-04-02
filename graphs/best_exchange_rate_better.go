package graphs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
)

// FindBestExchangeRatePath ...
func FindBestExchangeRatePath(ctx context.Context, sourceCurrency, targetCurrency string) (float64, []*GraphEdge, error) {
	// Initialize exchange client.
	cli := &ExchangeClient{
		&http.Client{
			Timeout: 30 * time.Second,
		},
		"fake-host",
	}

	// List exchange rates.
	exchangeRates, err := cli.ListExchangeRates(ctx)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to list exchange rates: %w", err)
	}

	// Build Graph.
	var (
		edges []*GraphEdge
		nodes = map[string]*GraphNode{}
	)
	for _, rate := range exchangeRates.Rates {
		if _, ok := nodes[rate.SourceCurrency]; !ok {
			nodes[rate.SourceCurrency] = &GraphNode{
				Currency: rate.SourceCurrency,
			}
		}

		if _, ok := nodes[rate.TargetCurrency]; !ok {
			nodes[rate.TargetCurrency] = &GraphNode{
				Currency: rate.TargetCurrency,
			}
		}

		nodeFro := nodes[rate.SourceCurrency]
		nodeTo := nodes[rate.TargetCurrency]

		ask, err := strconv.ParseFloat(rate.Ask, 64)
		if err != nil {
			// Log error.
			continue
		}

		bid, err := strconv.ParseFloat(rate.Bid, 64)
		if err != nil {
			// Log error.
			continue
		}

		edge := &GraphEdge{
			From: nodeFro,
			To:   nodeTo,
			Cost: normaliseCost(ask / bid),
		}

		nodeFro.Edges = append(nodeFro.Edges, edge)
		edges = append(edges, edge)
	}

	bestRate, path := BellmanFord(edges, sourceCurrency, targetCurrency)

	return math.Exp(math.Abs(bestRate)), path, nil
}

// ExchangeClient ...
type ExchangeClient struct {
	*http.Client
	host string
}

func (e *ExchangeClient) preparedURL(endpoint string) string {
	return fmt.Sprintf(e.host + endpoint) // use url.Parse here instead.
}

func (e *ExchangeClient) do(ctx context.Context, method, endpoint string, req, rsp interface{}) error {
	var body io.Reader
	if req != nil {
		rawBody, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("failed to marshal json body: %w", err)
		}
		body = bytes.NewReader(rawBody)
	}

	rawReq, err := http.NewRequestWithContext(ctx, method, e.preparedURL(endpoint), body)
	if err != nil {
	}

	rawRsp, err := e.Do(rawReq)
	if err != nil {
		return fmt.Errorf("failed to execute client request: %w", err)
	}
	defer rawRsp.Body.Close()

	rawRspBody, err := ioutil.ReadAll(rawReq.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(rawRspBody, rsp); err != nil {
		return fmt.Errorf("failed to unmarshal response body")
	}

	return nil
}

// ListExchangeRates ...
func (e *ExchangeClient) ListExchangeRates(ctx context.Context) (*ListExchangeRatesResponse, error) {
	var rsp = ListExchangeRatesResponse{}
	if err := e.do(ctx, http.MethodGet, "/rates", nil, rsp); err != nil {
		return nil, fmt.Errorf("failed to list exchange rates: %w", err)
	}

	return &rsp, nil
}

type ListExchangeRatesResponse struct {
	Rates []*ExchangeRateData `json:"rates"`
}

type ExchangeRateData struct {
	Bid            string `json:"bid"`
	Ask            string `json:"ask"`
	SourceCurrency string `json:"source_currency"`
	TargetCurrency string `json:"target_currency"`
	Rate           string `json:"rate"`
}

func normaliseCost(cost float64) float64 {
	return -math.Log(cost)
}

// GraphEdge ...
type GraphEdge struct {
	From, To *GraphNode
	Cost     float64
}

//GraphNode ...
type GraphNode struct {
	Edges    []*GraphEdge
	Currency string
}

// BellmanFord ...
func BellmanFord(edges []*GraphEdge, start, end string) (float64, []*GraphEdge) {
	var (
		currenciesToNode  = map[string]int{}
		nodesToCurrencies = map[int]string{}
	)

	// Build Graph.
	//
	// T -> O(v + e) -> O(n ** 2) as we have a complete graph.
	// S -> O(v)
	for _, edge := range edges {
		if _, ok := currenciesToNode[edge.From.Currency]; !ok {
			index := len(currenciesToNode)

			currenciesToNode[edge.From.Currency] = index
			nodesToCurrencies[index] = edge.From.Currency
		}

		if _, ok := currenciesToNode[edge.To.Currency]; !ok {
			index := len(currenciesToNode)

			currenciesToNode[edge.To.Currency] = index
			nodesToCurrencies[index] = edge.To.Currency
		}
	}

	var distances = make([]float64, len(currenciesToNode))
	for i := 0; i < len(currenciesToNode); i++ {
		distances[i] = math.MaxInt
	}

	startingNode := currenciesToNode[start]
	distances[startingNode] = 0

	var previous = make([]*GraphEdge, len(distances))

	// Relaxation of graph.
	//
	// T -> O(e ** 2) in a complete graph -> O(n ** 3)
	for k := 0; k < len(edges); k++ {
		for _, edge := range edges {
			toIndex := currenciesToNode[edge.To.Currency]
			fromIndex := currenciesToNode[edge.From.Currency]

			if distances[fromIndex] == math.MaxInt {
				distances[toIndex] = edge.Cost
				continue
			}

			if edge.Cost+distances[fromIndex] < distances[toIndex] {
				distances[toIndex] = edge.Cost + distances[fromIndex]
				previous[toIndex] = edge
			}
		}
	}

	// Propagate negative cycles.
	//
	// T -> O(e ** 2)
	for k := 0; k < len(edges); k++ {
		for _, edge := range edges {
			toIndex := currenciesToNode[edge.To.Currency]
			fromIndex := currenciesToNode[edge.From.Currency]

			if distances[fromIndex] == math.MaxInt && distances[toIndex] == math.MaxInt {
				continue
			}

			if distances[fromIndex] == math.MinInt {
				distances[toIndex] = math.MinInt
				continue
			}

			if edge.Cost+distances[fromIndex] < distances[toIndex] {
				distances[toIndex] = math.MinInt
			}
		}
	}

	var path []*GraphEdge

	// T -> O(v).
	var last = previous[startingNode]
	for last != nil {
		path = append(path, last)
		fromIndex := currenciesToNode[last.From.Currency]
		last = previous[fromIndex]
	}

	reversePath(path)

	// T -> O(v)
	return distances[startingNode], path
}

func reversePath(path []*GraphEdge) {
	var i, j = 0, len(path) - 1
	for i < j {
		path[i], path[j] = path[j], path[i]
	}
}
