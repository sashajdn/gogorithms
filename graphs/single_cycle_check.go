package graphs

// SingleCycleCheck Complexity:
// Time: O(n)
// Space: O(1)
func SingleCycleCheck(array []int) bool {
	currentIndex, itemsVisited := 0, 0

	for itemsVisited < len(array)-1 {
		if itemsVisited > 0 && itemsVisited < len(array) && currentIndex == 0 {
			return false
		}

		if itemsVisited == len(array) {
			return currentIndex == 0
		}

		itemsVisited++
		nextIndex := (array[currentIndex] + currentIndex) % len(array)
		currentIndex = nextIndex
		if currentIndex < 0 {
			currentIndex += len(array)
		}
	}
	return false
}
