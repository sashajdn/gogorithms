package algos

/*
Problem Statement:

Write a function that takes in an non-empty array of integers and returns
the maximum sum that can be obtained by summing up all of the integers in
a non-empty subarray of the input array. A subarray must only contain
adjacent numbers.

*/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


// Kadanes Algorithm
// Complexity:
// Time: O(n)
// Space: O(1)
func KadanesAlgorithm(array []int) int {
	sumSoFar := 0
	maxSoFar := array[0]

	for _, val := range array {
		sumSoFar = max(
			sumSoFar + val,
			val,
		)
		maxSoFar = max(
			sumSoFar,
			maxSoFar,
		)
	}
	return maxSoFar
} 
