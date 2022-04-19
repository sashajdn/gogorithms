package generators

import (
	"context"
	"sync"
)

// Iterator ...
type Iterator chan int

// Generator ...
type Generator func() Iterator

func (i Iterator) Next(ctx context.Context) (int, bool) {
	select {
	case v, ok := <-i:
		return v, ok
	case <-ctx.Done():
		return 0, false
	}
}

func (i Iterator) Generator(ctx context.Context, start, step, end int) Generator {
	return func() Iterator {
		return i.Range(ctx, start, step, end)
	}
}

// Range ...
func (i Iterator) Range(ctx context.Context, step, start, end int) Iterator {
	iterator := make(Iterator, 1)

	go func() {
		defer close(iterator)

		select {
		case <-ctx.Done():
			return
		default:
		}

		for j := start; j <= end; j += step {
			select {
			case iterator <- j:
			case <-ctx.Done():
				return
			}
		}
	}()

	return iterator
}

func Cycle(ctx context.Context, generator Generator) Iterator {
	cycledIterator := make(Iterator, 1)

	go func() {
		defer close(cycledIterator)

		for {
			currentIterator := generator()

			for v := range currentIterator {
				select {
				case cycledIterator <- v:
				case <-ctx.Done():
					return
				}
			}

		}
	}()

	return cycledIterator
}

func Negative(ctx context.Context, generator Generator) Iterator {
	iterator := generator()
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	negativeIterator := make(Iterator, 1)
	go func() {
		defer close(negativeIterator)
		for v := range iterator {
			select {
			case negativeIterator <- (-1 * v):
			case <-ctx.Done():
				return
			}
		}
	}()

	return negativeIterator
}

func Sink(ctx context.Context, iterators ...Iterator) Iterator {
	ctx, cancel := context.WithCancel(ctx)
	sink := make(Iterator, 1)

	var adjacentIterators = make([]Iterator, 0, len(iterators))
	for i := 0; i < len(iterators); i++ {
		adjacentIterators = append(adjacentIterators, make(Iterator, 1))
	}

	var wg sync.WaitGroup
	for i, iterator := range iterators {
		i, iterator := i, iterator
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			for v := range iterator {
				select {
				case adjacentIterators[i] <- v:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			for _, iterator := range adjacentIterators {
				v, ok := iterator.Next(ctx)
				if !ok {
					continue
				}

				select {
				case sink <- v:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		cancel()
		close(sink)
	}()

	return sink
}
