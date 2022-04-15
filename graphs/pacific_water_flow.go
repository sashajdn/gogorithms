package graphs

// PacificAtlantic ...
//
// T -> O(m * n)
// S -> O(m * n)
func PacificAtlantic(heights [][]int) [][]int {
	var (
		pacificQueue  = NewPALinkedList()
		atlanticQueue = NewPALinkedList()

		pacificAdj  = make([][]int, 0, len(heights))
		atlanticAdj = make([][]int, 0, len(heights))
	)

	for n := 0; n < len(heights); n++ {
		pacificAdj = append(pacificAdj, make([]int, len(heights[0])))
		atlanticAdj = append(atlanticAdj, make([]int, len(heights[0])))
	}

	for j := 0; j < len(heights); j++ {
		atlanticQueue.Add(&PANode{
			j: j,
			i: len(heights[0]) - 1,
		})
		atlanticAdj[j][len(heights[0])-1] = 1

		pacificQueue.Add(&PANode{
			j: j,
			i: 0,
		})
		pacificAdj[j][0] = 1
	}

	for i := 0; i < len(heights[0]); i++ {
		atlanticQueue.Add(&PANode{
			j: len(heights) - 1,
			i: i,
		})
		atlanticAdj[len(heights)-1][i] = 1

		pacificQueue.Add(&PANode{
			j: 0,
			i: i,
		})
		pacificAdj[0][i] = 1
	}

	bfs(heights, atlanticQueue, atlanticAdj)
	bfs(heights, pacificQueue, pacificAdj)

	var output = [][]int{}
	for j := 0; j < len(heights); j++ {
		for i := 0; i < len(heights[0]); i++ {
			if pacificAdj[j][i] == 2 && atlanticAdj[j][i] == 2 {
				output = append(output, []int{j, i})
			}
		}
	}

	return output
}

func bfs(heights [][]int, queue *PALinkedList, visited [][]int) {
	for !queue.IsEmpty() {
		next := queue.Pop()
		j, i := next.j, next.i

		if visited[j][i] > 0 {
			continue
		}

		for _, nn := range fetchValidNeighbours(heights, visited, j, i) {
			y, x := nn[0], nn[1]

			queue.Add(&PANode{
				j: y,
				i: x,
			})

			visited[j][i] = 2
		}
	}
}

func fetchValidNeighbours(heights, visited [][]int, j, i int) [][]int {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var validNeighbours = make([][]int, 0, 4)
	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[1]

		if dj < 0 || dj >= len(heights) {
			continue
		}

		if di < 0 || di >= len(heights[0]) {
			continue
		}

		if visited[dj][di] > 0 {
			continue
		}
		visited[dj][di] = 1

		if heights[dj][di] < heights[j][i] {
			continue
		}

		validNeighbours = append(validNeighbours, []int{dj, di})
	}

	return validNeighbours
}

type PANode struct {
	Previous, Next *PANode
	i, j           int
}

type PALinkedList struct {
	Head   *PANode
	Tail   *PANode
	length int
}

func (l *PALinkedList) Len() int {
	return l.length
}

func NewPALinkedList() *PALinkedList {
	head := &PANode{}
	tail := &PANode{}

	head.Next = tail
	tail.Previous = head

	return &PALinkedList{
		Head: head,
		Tail: tail,
	}
}

func (l *PALinkedList) Pop() *PANode {
	if l.Tail.Previous == l.Head {
		return nil
	}

	previous := l.Tail.Previous
	l.Tail.Previous = l.Tail.Previous.Previous
	l.Tail.Previous.Next = nil

	if l.Tail.Previous == l.Head {
		l.Head.Next = l.Tail
	}

	l.length--
	return previous
}

func (l *PALinkedList) Add(node *PANode) {
	currentHead := l.Head
	l.Head = node

	node.Next = currentHead
	currentHead.Previous = node

	l.length++
}

func (l *PALinkedList) IsEmpty() bool {
	return l.Head.Next == l.Tail
}
