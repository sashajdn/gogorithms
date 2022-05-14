package graphs

import (
	"container/heap"
	"math"
)

// 0 -> n - 1 node
// SSSP
// Dijkstas (Priority Queue): m -> m(e)
// Peek, Pop, Push: O(1), O(log(n))
// O(log(n))
//
// costs []int{0, +inf, +inf, +inf, ..., +inf}
// 0 -> 2 cost 10
//
// costs[from] + 10 < costs[to] -> Push Edge, don't decrement discounts
// costs[from] + (10 / 2) < costs[to] -> Push Edge, do decrement discount

// MinCostToReachCityWithDiscount ...
func MinCostToReachCityWithDiscount(n int, highways [][]int, discounts int) int {
	var graph = map[int][][]int{}
	for _, highway := range highways {
		from, to, cost := highway[0], highway[1], highway[2]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, cost})

		if _, ok := graph[to]; !ok {
			graph[to] = [][]int{}
		}
		graph[to] = append(graph[to], []int{from, cost})
	}

	var pq = &DiscountsPQ{}
	heap.Push(pq, &DiscountsPQItem{
		Index:         0,
		DiscountsUsed: discounts,
		Cost:          0,
	})

	var costs = make([][]int, 0, n)
	for j := 0; j < discounts; j++ {
		var discountedCosts = make([]int, n)
		for i := 1; i < len(costs); i++ {
			discountedCosts[i] = math.MaxInt
		}
		costs = append(costs, discountedCosts)
	}

	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*DiscountsPQItem)

		for _, edge := range graph[from.Index] {
			to, cost := edge[0], edge[1]

			if from.Cost+cost < costs[to][from.DiscountsUsed] {
				heap.Push(pq, &DiscountsPQItem{
					Index:         to,
					Cost:          cost + from.Cost,
					DiscountsUsed: from.DiscountsUsed,
				})
			}

			if from.Cost+(cost/2) < costs[to][from.DiscountsUsed] && from.DiscountsUsed < discounts {
				heap.Push(pq, &DiscountsPQItem{
					Index:         to,
					Cost:          cost + from.Cost,
					DiscountsUsed: from.DiscountsUsed + 1,
				})
			}
		}
	}

	if costs[n-1][discounts-1] == math.MaxInt {
		return -1
	}

	return costs[n-1][discounts-1]
}

type DiscountsPQItem struct {
	DiscountsUsed, Index, Cost int
}

type DiscountsPQ []*DiscountsPQItem

func (d DiscountsPQ) Len() int           { return len(d) }
func (d DiscountsPQ) Less(i, j int) bool { return d[i].Cost < d[j].Cost }
func (d DiscountsPQ) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d *DiscountsPQ) Push(value interface{}) {
	v := value.(*DiscountsPQItem)
	*d = append(*d, v)
}
func (d *DiscountsPQ) Pop() interface{} {
	v := (*d)[d.Len()-1]
	*d = (*d)[:d.Len()-1]
	return v
}
