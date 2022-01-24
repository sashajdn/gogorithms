package graphs

// NumberOfWaysToTraverseGraph accepts two params: width & heigth that serves as the starting point at the top left,
// assuming that travel can only take one of two ways: right & down, this returns the total number of
// ways to traverse a graph.
//
// T -> O(nm)
// S -> O(nm)
func NumberOfWaysToTraverseGraph_Dynamic(width, height int) int {
	// Initialization.
	var wh = make([][]int, 0, height)

	for j := 0; j < height; j++ {
		wh = append(wh, make([]int, width))
	}

	// Set all elements in the first row to 1.
	for i := 0; i < width; i++ {
		wh[0][i] = 1
	}

	// Set all elements in the first column to 1.
	for j := 0; j < height; j++ {
		wh[j][0] = 1
	}

	// Add the previous row & column & set.
	for j := 1; j < height; j++ {
		for i := 1; i < width; i++ {
			wh[j][i] = wh[j-1][i] + wh[j][i-1]
		}
	}

	return wh[height-1][width-1]
}
