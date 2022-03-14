package graphs

type NAryTree struct {
	Value    int
	Children []*NAryTree
}

type NodeInfo struct {
	Height   int
	Diameter int
}

// DiameterNAryTree ...
//
// T -> O(v + e) where `v` is the vertices of the tree & `e` is the edges. Here since we maintain only the top two
// 	         greatest heights of a given nodes children, we have constant time - unlike if we sorted or used a heap.
// S -> O(v + e)
func DiameterNAryTree(root *NAryTree) int {
	info := findDiameter(root)
	return info.Diameter
}

func findDiameter(node *NAryTree) *NodeInfo {
	if node == nil {
		return &NodeInfo{}
	}

	var a, b = &NodeInfo{}, &NodeInfo{}
	for _, child := range node.Children {
		childInfo := findDiameter(child)
		parentHeight := childInfo.Height + 1
		switch {
		case parentHeight > a.Height:
			b = a
			a = childInfo
			continue
		case parentHeight > b.Height:
			b = childInfo
			continue
		}
	}

	summedHeight := a.Height + b.Height
	if summedHeight > a.Diameter && summedHeight > b.Diameter {
		return &NodeInfo{
			Height:   max(a.Height, b.Height) + 1,
			Diameter: summedHeight,
		}
	}

	return &NodeInfo{
		Height:   max(a.Height, b.Height) + 1,
		Diameter: max(a.Diameter, b.Diameter),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
