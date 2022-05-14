package graphs

import (
	"container/heap"
	"math"
)

// MinCostReachCityWithDiscount_2D ...
//
// T -> O((v + e) * log(v))
// S -> O((v + e) + (v * d))
func MinCostReachCityWithDiscount_2D(n int, highways [][]int, discounts int) int {
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

	var costs = make([][]int, 0, discounts+1)
	for i := 0; i < discounts+1; i++ {
		var discountsCosts = make([]int, n)
		for i = 0; i < n; i++ {
			discountsCosts[i] = math.MaxInt
		}
		costs = append(costs, discountsCosts)
	}
	costs[discounts][0] = 0

	var pq = &PQ{}
	heap.Push(pq, &DiscountsPQItem{
		Index:         0,
		Cost:          0,
		DiscountsLeft: discounts,
	})

	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*DiscountsPQItem)

		if from.Index == n-1 {
			return from.Cost
		}

		for _, edge := range graph[from.Index] {
			to, cost := edge[0], edge[1]

			if cost+from.Cost < costs[from.DiscountsLeft][to] {
				costs[from.DiscountsLeft][to] = cost + from.Cost

				heap.Push(pq, &DiscountsPQItem{
					Index:         to,
					Cost:          cost + from.Cost,
					DiscountsLeft: from.DiscountsLeft,
				})
			}

			if from.DiscountsLeft > 0 && (cost/2)+from.Cost < costs[from.DiscountsLeft-1][to] {
				costs[from.DiscountsLeft-1][to] = cost + from.Cost

				heap.Push(pq, &DiscountsPQItem{
					Index:         to,
					Cost:          cost + from.Cost,
					DiscountsLeft: from.DiscountsLeft - 1,
				})
			}
		}
	}

	return -1
}

// MinCostReachCityWithDiscount_1D ...
//
// T -> O((v + e) * log(v))
// S -> O(v + e)
func MinCostReachCityWithDiscount_1D(n int, highways [][]int, discounts int) int {
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

	var pq = &PQ{}
	heap.Push(pq, &DiscountsPQItem{
		Index:         0,
		Cost:          0,
		DiscountsLeft: discounts,
	})

	var visitedDiscounts = make(map[int]int, n)
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*DiscountsPQItem)

		if v, ok := visitedDiscounts[from.Index]; ok && from.DiscountsLeft <= v {
			continue
		}
		visitedDiscounts[from.Index] = from.DiscountsLeft

		if from.Index == n-1 {
			return from.Cost
		}

		for _, edge := range graph[from.Index] {
			to, cost := edge[0], edge[1]

			if from.DiscountsLeft > 0 {
				heap.Push(pq, &DiscountsPQItem{
					Cost:          (cost / 2) + from.Cost,
					DiscountsLeft: from.DiscountsLeft - 1,
					Index:         to,
				})
			}

			heap.Push(pq, &DiscountsPQItem{
				Cost:          cost + from.Cost,
				DiscountsLeft: from.DiscountsLeft,
				Index:         to,
			})
		}
	}

	return -1
}

type DiscountsPQItem struct {
	Index, Cost, DiscountsLeft int
}

type DiscountsPQ []*DiscountsPQItem

func (d DiscountsPQ) Len() int           { return len(d) }
func (d DiscountsPQ) Less(i, j int) bool { return d[i].Cost < d[j].Cost }
func (d DiscountsPQ) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d *DiscountsPQ) Pop() interface{} {
	v := (*d)[d.Len()-1]
	*d = (*d)[:d.Len()-1]
	return v
}
func (d *DiscountsPQ) Push(value interface{}) {
	vv := value.(*DiscountsPQItem)
	*d = append(*d, vv)
}
