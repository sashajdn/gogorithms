package iterators

import "context"

func StatefulIteratorFromArrays(arrays ...[]int) *StatefulIterator {
	var iterators = make([][]int, 0, len(arrays))
	for _, array := range arrays {
		iterators = append(iterators, array)
	}

	return &StatefulIterator{
		sources:   iterators,
		exhausted: make(map[int]struct{}),
		cursor:    0,
		last:      make([]int, len(arrays)),
	}
}

type StatefulIterator struct {
	sources   [][]int
	exhausted map[int]struct{}
	cursor    int
	last      []int
	cycling   bool
}

func (s *StatefulIterator) Range() <-chan int {
	ch := make(chan int, 1)

	go func() {
		for len(s.exhausted) < len(s.sources) {
			if _, ok := s.exhausted[s.cursor]; ok {
				s.cursor = (s.cursor + 1) % len(s.sources)
				continue
			}

			next := s.last[s.cursor]
			if len(s.sources[s.cursor]) != 0 && next < len(s.sources[s.cursor]) {
				ch <- s.sources[s.cursor][next]

				s.last[s.cursor]++
				s.cursor = (s.cursor + 1) % len(s.sources)
				continue
			}

			s.exhausted[s.cursor] = struct{}{}
			s.cursor = (s.cursor + 1) % len(s.sources)
		}

		close(ch)
	}()

	return ch
}

func (s *StatefulIterator) Cycle(ctx context.Context) <-chan int {
	ch := make(chan int, 1)
	s.cycling = true

	go func() {
		defer func() {
			s.Reset()
			close(ch)
			s.cycling = false
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			for len(s.exhausted) < len(s.sources) {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if _, ok := s.exhausted[s.cursor]; ok {
					s.cursor = (s.cursor + 1) % len(s.sources)
					continue
				}

				next := s.last[s.cursor]
				if len(s.sources[s.cursor]) != 0 && next < len(s.sources[s.cursor]) {
					ch <- s.sources[s.cursor][next]

					s.last[s.cursor]++
					s.cursor = (s.cursor + 1) % len(s.sources)
					continue
				}

				s.exhausted[s.cursor] = struct{}{}
				s.cursor = (s.cursor + 1) % len(s.sources)
			}

			s.Reset()
		}
	}()

	return ch
}

func (s *StatefulIterator) Reset() {
	s.cursor = 0
	s.last = make([]int, len(s.sources))
	s.exhausted = map[int]struct{}{}
}

func (s *StatefulIterator) Next() (int, bool) {
	for len(s.exhausted) < len(s.sources) {
		if _, ok := s.exhausted[s.cursor]; ok {
			s.cursor = (s.cursor + 1) % len(s.sources)
			continue
		}

		next := s.last[s.cursor]
		if len(s.sources[s.cursor]) != 0 && next < len(s.sources[s.cursor]) {
			v := s.sources[s.cursor][next]
			s.last[s.cursor]++
			s.cursor = (s.cursor + 1) % len(s.sources)
			return v, true
		}

		s.cursor = (s.cursor + 1) % len(s.sources)
	}

	return 0, false
}

func (s *StatefulIterator) HasNext() bool {
	return (len(s.exhausted) < len(s.sources)) || s.cycling
}
