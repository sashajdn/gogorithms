package clocks

import (
	"fmt"
	"sync"
)

// VectorClock is a vector clock
type VectorClock struct {
	sync.Mutex
	id        int
	timestamp []int
}

// New constructs a new Vector Clock & returns a pointer to it
func New(id int, numberOfNodes int) *VectorClock {
	return &VectorClock{
		id:        id,
		timestamp: make([]int, numberOfNodes),
	}
}

// Increment increments the timestamp for the given node
func (v *VectorClock) Increment() error {
	if v.id > len(v.timestamp) {
		return fmt.Errorf("Not enough nodes are initialized")
	}
	return v.IncrementOther(v.id)
}

func (v *VectorClock) IncrementOther(id int) error {
	if id > len(v.timestamp) {
		return fmt.Errorf("Not enough nodes are initialized")
	}

	newTimestamp := make([]int, len(v.timestamp))
	copy(newTimestamp, v.timestamp)
	newTimestamp[id]++

	return v.Merge(newTimestamp)
}

// Merge merges two timestamps and sets as the current timestamp
func (v *VectorClock) Merge(newTimestamp []int) error {
	v.Lock()
	defer v.Unlock()
	if len(newTimestamp) != len(v.timestamp) {
		return fmt.Errorf("Cannot compare timestamps; differing sizes")
	}
	for i, val := range newTimestamp {
		v.timestamp[i] = max(v.timestamp[i], val)
	}
	return nil
}

// Equal checks if two timestamps occur at the same time.
func (v *VectorClock) Equal(timestamp []int) (bool, error) {
	if len(v.timestamp) != len(timestamp) {
		return false, fmt.Errorf("Cannot compare timestamps; differing sizes")
	}
	for i, val := range timestamp {
		if v.timestamp[i] != val {
			return false, nil
		}
	}
	return true, nil
}

// BeforeEqual checks if an event occurs after another event by comparing the timestamps, or if they are equal
func (v *VectorClock) BeforeEqual(timestamp []int) (bool, error) {
	if len(v.timestamp) != len(timestamp) {
		return false, fmt.Errorf("Cannot compare timestamps; differing sizes")
	}
	for i, val := range timestamp {
		if v.timestamp[i] > val {
			return false, nil
		}
	}
	return true, nil
}

// AfterEqual checks if an event occurs after another event by comparing the timestamps, or if they are equal
func (v *VectorClock) AfterEqual(timestamp []int) (bool, error) {
	if len(v.timestamp) != len(timestamp) {
		return false, fmt.Errorf("Cannot compare timestamps; differing sizes")
	}
	for i, val := range timestamp {
		if v.timestamp[i] < val {
			return false, nil
		}
	}
	return true, nil
}

// Before checks if an event occurs after another event by comparing the timestamps given they are not equal
func (v *VectorClock) Before(timestamp []int) (bool, error) {
	is, err := v.AfterEqual(timestamp)
	return !is, err
}

// After checks if an event occurs after another event by comparing the timestamps given they are not equal
func (v *VectorClock) After(timestamp []int) (bool, error) {
	is, err := v.BeforeEqual(timestamp)
	return !is, err
}

// IsConcurrent checks if two events are concurrent by comparing the timestamp of each event.
// Two events are concurrent is t does NOT occur before or after t', and t is not equal to t'
func (v *VectorClock) IsConcurrent(timestamp []int) (bool, error) {
	is, err := v.BeforeEqual(timestamp)
	if err != nil {
		return false, err
	}
	if is {
		return false, err
	}
	is, err = v.AfterEqual(timestamp)
	if err != nil {
		return false, err
	}
	if is {
		return false, err
	}
	return true, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
