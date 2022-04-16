package intervals

import "sort"

// CanAttendMeetings ...
//
// T -> O(nlog(n)) where `n` is the number of meetings.
// S -> O(1) if we do the sorting in place.
func CanAttendMeetings(meetings [][]int) bool {
	if len(meetings) < 2 {
		return true
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	var previous = meetings[0]
	for i := 1; i < len(meetings); i++ {
		current := meetings[i]

		if current[0] < previous[1] {
			return false
		}

		previous[1] = max(previous[1], current[1])
	}

	return true
}
