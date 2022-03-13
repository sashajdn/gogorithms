package binarytree

import "sort"

type NodeInfo struct {
	Value int
	Row   int
}

type NodeInfoList []*NodeInfo

func (n NodeInfoList) Len() int { return len(n) }
func (n NodeInfoList) Less(a, b int) bool {
	if n[a].Row == n[b].Row {
		return n[a].Value <= n[b].Value
	}

	return n[a].Row <= n[b].Row
}
func (n NodeInfoList) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

// VerticalTraversal_Hard
//
// T -> O(w * h)
// S -> O(max(w * h), n)
func VerticalTraversal_Hard(root *TreeNode) [][]int {
	var (
		mapping   = map[int]*NodeInfoList{}
		traversal = [][]int{}
	)

	minBound, maxBound := dfs(root, 0, 0, 0, 0, mapping)

	for i := minBound; i <= maxBound; i++ {
		colwise, ok := mapping[i]
		if !ok {
			continue
		}

		sort.Sort(colwise)

		var entry []int
		for _, v := range *colwise {
			entry = append(entry, v.Value)
		}

		traversal = append(traversal, entry)
	}

	return traversal
}

func dfs(node *TreeNode, col, row, min, max int, mapping map[int]*NodeInfoList) (int, int) {
	if node == nil {
		return minF(col, min), maxF(col, max)
	}

	if _, ok := mapping[col]; !ok {
		mapping[col] = &NodeInfoList{}
	}

	*mapping[col] = append(*mapping[col], &NodeInfo{
		Value: node.Val,
		Row:   row,
	})

	lmn, lmx := dfs(node.Left, col-1, row+1, min, max, mapping)
	rmn, rmx := dfs(node.Right, col+1, row+1, min, max, mapping)
	return minF(lmn, rmn), maxF(lmx, rmx)

}

func minF(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}
