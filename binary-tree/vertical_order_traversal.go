package binarytree

import "sort"

// VerticalOrder ...
func VerticalOrder(root *TreeNode) [][]int {
	var (
		output [][]int
		hm     = make(map[int]*[][]int)
	)

	min, max := collect(root, 0, 0, hm)

	r := max - min
	for i := 1; i < r; i++ {
		rows := *hm[min+i]
		sort.Slice(rows, func(i, j int) bool {
			return rows[i][1] > rows[j][0]
		})

		var rowValues []int
		for _, row := range rows {
			rowValues = append(rowValues, row[0])
		}

		output = append(output, rowValues)
	}

	return output
}

func collect(root *TreeNode, column, row int, hm map[int]*[][]int) (int, int) {
	if root == nil {
		return column, column
	}

	v, ok := hm[column]
	switch {
	case ok:
		*v = append(*v, []int{root.Val, row})
	default:
		hm[column] = &[][]int{{root.Val, row}}
	}

	lmin, lmax := collect(root.Left, column-1, row+1, hm)
	rmin, rmax := collect(root.Right, column+1, row+1, hm)

	return min(lmin, rmin), max(lmax, rmax)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
