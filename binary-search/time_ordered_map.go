package binarysearch

// TimeOrderedMap ...
type TimeOrderedMap struct {
	store map[string][]*TimeOrderedNode
}

// TimeOrderedNode ...
type TimeOrderedNode struct {
	Value     string
	Timestamp int
}

// Set ...
//
// T -> O(1)
// S -> O(1)
func (t *TimeOrderedMap) Set(key, value string, timestamp int) {
	if _, ok := t.store[key]; !ok {
		t.store[key] = []*TimeOrderedNode{}
	}

	t.store[key] = append(t.store[key], &TimeOrderedNode{
		Value:     value,
		Timestamp: timestamp,
	})
}

// Get ...
//
// T -> O(log(n)) where `n` is the total number of records in the map in the worst case.
// S -> O(n)
func (t *TimeOrderedMap) Get(key string, timestamp int) string {
	workingSet, ok := t.store[key]
	if !ok {
		return ""
	}
	if len(workingSet) == 0 {
		return ""
	}

	var left, right = 0, len(workingSet) - 1
	for left < right {
		mid := (left + right) / 2

		var workingValue = workingSet[mid]
		if workingValue.Timestamp == timestamp {
			return workingValue.Value
		}

		if timestamp < workingValue.Timestamp {
			right = mid - 1
			continue
		}

		left = mid + 1
	}

	var workingValue = workingSet[left]
	if left == 0 && workingValue.Timestamp > timestamp {
		return ""
	}

	if workingValue.Timestamp > timestamp {
		return workingSet[left-1].Value
	}

	return workingValue.Value
}
