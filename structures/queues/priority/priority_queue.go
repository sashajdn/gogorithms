package priority

type PriorityQueue []int

func (p PriorityQueue) Len() int { return len(p) }
func (p PriorityQueue) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *PriorityQueue) Push(value interface{}) {
	vv, _ := value.(int)
	*p = append(*p, vv)
}

func (p *PriorityQueue) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
