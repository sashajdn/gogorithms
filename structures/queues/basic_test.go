package queues

import "testing"

func TestBasicQueue(t *testing.T) {
	t.Parallel()
	q := NewBasicQueue(-1)

	input := []interface{}{1, 2, 3}
	q.EnqueueMany(input)
	assertSlice(t, input, q.q, slicesEqual)
	assert(t, 3, q.Size())

	v, ok := q.Pop()
	assert(t, 1, v)
	assert(t, true, ok)
	assert(t, 2, q.Size())
	assertSlice(t, input[1:], q.q, slicesEqual)

	v, ok = q.Peek()
	assert(t, 2, v)
	assert(t, true, ok)
	assert(t, 2, q.Size())
	assertSlice(t, input[1:], q.q, slicesEqual)
}
