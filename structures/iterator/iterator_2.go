package iterator

// Generator ...
type Generator interface {
	Next()
	HasNext()
	SetBroker(chan int)
}

// IteratorWithGenerators ...
type IteratorWithGenerators struct {
	Generators []Generator
}
