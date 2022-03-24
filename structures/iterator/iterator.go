package iterator

import (
	"context"
)

type IntIterator interface {
	Next() int
	HasNext() int
	AsyncNext() <-chan int
	Reset(start, end, step int)
	SetFilter(func(v int) bool)
}

type Iterator struct {
	it               chan int
	start, end, step int
	filter           func(v int) bool
	ctx              context.Context
	cancel           func()
}

type IteratorOpt func(it Iterator)

func NewIterator(opts ...IteratorOpt) Iterator {
	it := Iterator{
		it:    make(chan int, 1),
		start: 0,
		end:   -1,
		step:  1,
	}

	for _, o := range opts {
		o(it)
	}

	it.ctx, it.cancel = context.WithCancel(context.Background())
	go it.iterate()

	return it
}

func (i *Iterator) iterate() {
	defer close(i.it)

	switch {
	case i.end == -1:
		var v = i.start
		for {
			select {
			case <-i.ctx.Done():
			default:
				return
			}

			if !i.filter(v) {
				v += i.step
				continue
			}

			i.it <- v
		}
	default:
		for j := i.start; j <= i.end; j += i.step {
			select {
			case <-i.ctx.Done():
			default:
				return
			}

			i.it <- j
		}
	}
}

func (i *Iterator) SetFilter(filter func(v int) bool) {
	i.filter = filter
}

func (i *Iterator) Reset(start, end, step int) {
	// Cancel.
	i.cancel()

	// Drain.
	for k := 0; k < 2; k++ {
		select {
		case <-i.it:
		}
	}

	// Reset iterator.
	i.ctx, i.cancel = context.WithCancel(context.Background())
	i.start, i.end, i.step = start, end, step

	// Start iterator.
	go i.iterate()
}

func (i *Iterator) Next() int {
	select {
	case v := <-i.it:
		return v
	default:
		return 0 // default value to send.
	}
}

func (i *Iterator) HasNext() bool {
	if i.end == -1 || i.start < i.end {
		return true
	}

	return false
}

func (i *Iterator) NextAsync(ctx context.Context) <-chan int {
	tmpCh := make(chan int)

	go func() {
		select {
		case <-ctx.Done():
			return
		case v := <-i.it:
			tmpCh <- v
		}
	}()

	return tmpCh
}

func IteratorOptSetStart(it Iterator, start int) {
	it.start = start
}
func IteratorOptSetEnd(it Iterator, end int) {
	it.end = end
}
func IteratorOptSetStep(it Iterator, step int) {
	it.step = step
}
