package algorithms

func GenerateDocument(characters, document string) bool {
	if document == "" {
		return true
	}
	counter := Counter{
		m: map[rune]int{},
	}
	for _, r := range document {
		counter.Incr(r)
	}
	for _, r := range characters {
		counter.Decr(r)
	}
	return counter.Empty()
}

type Counter struct {
	m map[rune]int
}

func (c *Counter) Incr(r rune) {
	c.m[r]++
}

func (c *Counter) Decr(r rune) {
	if _, ok := c.m[r]; !ok {
		return
	}
	if c.m[r] == 1 {
		delete(c.m, r)
		return
	}
	c.m[r]--
}

func (c *Counter) Empty() bool {
	return len(c.m) == 0
}
