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

// NumberOfWaysToTraverseGraph_Recursive implements the above but recusively.
//
// T -> O(2 ** (n+m))
// S -> O(n + m)
func NumberOfWaysToTraverseGraph_Recursive(width, height int) int {
	if width == 1 || height == 1 {
		return 1
	}

	return NumberOfWaysToTraverseGraph_Recursive(width-1, height) + NumberOfWaysToTraverseGraph_Recursive(width, height-1)
}

// NumberOfWaysToTraverseGraph_Factorial ...
//
// T -> O(nm)
// S -> O(1)
func NumberOfWaysToTraverseGraph_Factorial(width, height int) int {
	// Number of perms of n possible options -> (sum(0 -> n))! / product(i!)0->i
	return factorial(width+height-2) / (factorial(width-1) * factorial(height-1))
}

func factorial(n int) int {
	if n < 2 {
		return 1
	}

	return n * factorial(n-1)
}
