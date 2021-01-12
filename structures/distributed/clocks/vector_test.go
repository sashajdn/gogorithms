package clocks

import "testing"

func TestVectorClock(t *testing.T) {
	t.Parallel()

	numberOfNodes := 10

	vc1 := New(2, numberOfNodes)
	vc2 := New(6, numberOfNodes)

	vc1.Increment()
	assert(t, 1, vc1.timestamp[2], func(a, b interface{}) bool {
		return a == b
	})

	vc2.Increment()
	vc1.Merge(vc2.timestamp)
	assert(t, 1, vc1.timestamp[2], func(a, b interface{}) bool {
		return a == b
	})
	assert(t, 1, vc1.timestamp[6], func(a, b interface{}) bool {
		return a == b
	})

	vc1.Increment()

	res, err := vc1.Equal(vc2.timestamp)
	assertNoError(t, err)
	assert(t, false, res, func(a, b interface{}) bool {
		return a == b
	})

	res, err = vc1.Before(vc2.timestamp)
	assertNoError(t, err)
	assert(t, false, res, func(a, b interface{}) bool {
		return a == b
	})

	res, err = vc1.BeforeEqual(vc2.timestamp)
	assertNoError(t, err)
	assert(t, false, res, func(a, b interface{}) bool {
		return a == b
	})

	res, err = vc1.After(vc2.timestamp)
	assertNoError(t, err)
	assert(t, true, res, func(a, b interface{}) bool {
		return a == b
	})

	res, err = vc1.AfterEqual(vc2.timestamp)
	assertNoError(t, err)
	assert(t, true, res, func(a, b interface{}) bool {
		return a == b
	})
}

func assert(t *testing.T, expected, value interface{}, equalFunc func(a, b interface{}) bool) {
	if !equalFunc(expected, value) {
		t.Fatalf("expectedk -> %v, got -> %v", expected, value)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Not expecting Error: %v", err)
	}
}
