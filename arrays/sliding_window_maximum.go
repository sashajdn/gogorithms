package arrays

import (
	"math"
)

// SlidingWindowMaximum_BruteForce ...
//
// T -> O(nk)
// S -> O(k)
func SlidingWindowMaximum_BruteForce(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	if len(nums) == 1 || k == 1 {
		return nums
	}

	var (
		output        = make([]int, 0, len(nums)-k+1)
		slidingWindow = make([]int, 0, k)
	)

	var maxInWindow = math.MinInt
	for i := 0; i < k; i++ {
		slidingWindow = append(slidingWindow, nums[i])
		if nums[i] > maxInWindow {
			maxInWindow = nums[i]
		}
	}
	output = append(output, maxInWindow)

	// T -> O(n * k)
	// S -> O(k)
	for j := k; j < len(nums); j++ {
		var newMax = math.MinInt
		for i := 1; i < k; i++ {
			newMax = max(newMax, slidingWindow[i])
			slidingWindow[i-1], slidingWindow[i] = slidingWindow[i], slidingWindow[i-1]
		}

		newMax = max(newMax, nums[j])
		slidingWindow[k-1] = nums[j]

		output = append(output, newMax)
	}

	return output
}

type DequeNode struct {
	Next, Previous *DequeNode
	Index          int
	ID             string
}

type Deque struct {
	Head, Tail *DequeNode
	length     int
}

func NewDeque() *Deque {
	head := &DequeNode{
		Index: -1,
		ID:    "HEAD",
	}
	tail := &DequeNode{
		Index: -1,
		ID:    "TAIL",
	}

	head.Next = tail
	tail.Previous = head

	return &Deque{
		Head: head,
		Tail: tail,
	}
}

// Pop ...
func (d *Deque) Pop() *DequeNode {
	if d.Tail.Previous.ID == "HEAD" {
		return nil
	}

	previous := d.Tail.Previous

	previousPrevious := d.Tail.Previous.Previous
	previousPrevious.Next = d.Tail
	d.Tail.Previous = previousPrevious

	previous.Previous = nil
	previous.Next = nil

	d.length--

	return previous
}

// PopLeft ...
func (d *Deque) PopLeft() *DequeNode {
	if d.Head.Next.ID == "TAIL" {
		return nil
	}

	next := d.Head.Next

	nextNext := d.Head.Next.Next
	nextNext.Previous = d.Head
	d.Head.Next = nextNext

	next.Next = nil
	next.Previous = nil

	d.length--

	return next
}

// Append ...
func (d *Deque) Append(index int) {
	previous := d.Tail.Previous

	node := &DequeNode{
		Index:    index,
		Next:     d.Tail,
		Previous: previous,
	}

	previous.Next = node
	d.Tail.Previous = node

	d.length++
}

func (d *Deque) PeekFirst() (int, bool) {
	if d.Head.Next == d.Tail {
		return 0, false
	}

	return d.Head.Next.Index, true
}

func (d *Deque) PeekLast() (int, bool) {
	if d.Tail.Previous == d.Head {
		return 0, false
	}

	return d.Tail.Previous.Index, true
}

func (d *Deque) Len() int {
	return d.length
}

// SlidingWindowMaximum_Deque ...
//
// T -> O(n)
// S -> O(n)
func SlidingWindowMaximum_Deque(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	var (
		slidingWindow = NewDeque()
		output        = make([]int, 0, len(nums)-k+1)
	)

	var maxSoFar = math.MinInt
	for i := 0; i < k; i++ {
		for slidingWindow.Len() > 0 {
			lastIndex, ok := slidingWindow.PeekLast()
			if !ok {
				continue
			}

			if nums[i] > nums[lastIndex] {
				_ = slidingWindow.Pop()
				continue
			}

			break
		}

		slidingWindow.Append(i)
		maxSoFar = max(maxSoFar, nums[i])
	}
	output = append(output, maxSoFar)

	for j := k; j < len(nums); j++ {
		firstIndex, ok := slidingWindow.PeekFirst()
		if !ok {
			continue
		}

		if firstIndex <= j-k {
			_ = slidingWindow.PopLeft()
		}

		for slidingWindow.Len() > 0 {
			lastIndex, ok := slidingWindow.PeekLast()
			if !ok {
				continue
			}

			if nums[j] > nums[lastIndex] {
				_ = slidingWindow.Pop()
				continue
			}

			break
		}
		slidingWindow.Append(j)

		firstIndex, ok = slidingWindow.PeekFirst()
		if !ok {
			break
		}

		output = append(output, nums[firstIndex])
	}

	return output
}

// SlidingWindowMaximum_Dynamic ...
// T -> O(n)
// S -> O(n)
func SlidingWindowMaximum_Dynamic(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	if len(nums) == 1 || k == 1 {
		return nums
	}

	var (
		left, right   = make([]int, len(nums)), make([]int, len(nums))
		leftMaxSoFar  = math.MinInt
		rightMaxSoFar = math.MinInt
	)
	for i := 0; i < len(nums); i++ {
		leftMaxSoFar = max(leftMaxSoFar, nums[i])
		left[i] = leftMaxSoFar

		if (i+1)%k == 0 {
			leftMaxSoFar = math.MinInt
		}

		rightIndex := len(nums) - i - 1

		rightMaxSoFar = max(rightMaxSoFar, nums[rightIndex])
		right[rightIndex] = rightMaxSoFar

		if rightIndex%k == 0 {
			rightMaxSoFar = math.MinInt
		}
	}

	var output = make([]int, 0, len(nums)-k+1)
	for i := k - 1; i < len(nums); i++ {
		output = append(output, max(left[i], right[i-k+1]))
	}

	return output
}
