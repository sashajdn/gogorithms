package intervals

import (
	"fmt"
	"strconv"
	"strings"
)

type StringMeeting struct {
	Start string
	End   string
}

// [11:00, 12:00], [12:55, 14:00]
// 1. merge two calanders t -> O(n + m), s(n + m)
// 2. > max(start(n), start(m)), min((end(n), end(m))) // Start
// 3 meetingDuration 60 mins [...[1, 2], [ 3:30, 4]...] -> [[2:00, 3:30]], [[[2:00, 3,00], [2:01, 3:01]]
func CalendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	var availableSlots []StringMeeting
	return availableSlots
}

func mergeCalanders(a, b []StringMeeting) []StringMeeting {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	var mergedCalanders []StringMeeting
	l, r := 0, 0
	switch {
	case MilitaryTime(a[l].Start).LessThan(MilitaryTime(b[r].Start)):
		mergedCalanders = append(mergedCalanders, a[l])
		l++
	case MilitaryTime(b[r].Start).LessThan(MilitaryTime(a[l].Start)):
		mergedCalanders = append(mergedCalanders, b[r])
		r++
	}

	for l > len(a) && r > len(b) {
		switch {
		case MilitaryTime(a[l].Start).LessThan(MilitaryTime(b[r].Start)):
			if MilitaryTime(a[l].Start).LessThan(MilitaryTime(mergedCalanders[len(mergedCalanders)-1].End)) {
				mergedCalanders[len(mergedCalanders)-1].End = max(mergedCalanders[len(mergedCalanders)-1].End, a[l].End)
			} else {
				mergedCalanders = append(mergedCalanders, a[l])
			}
			l++
		case MilitaryTime(b[r].Start).LessThan(MilitaryTime(a[l].Start)):
			if MilitaryTime(b[r].Start).LessThan(MilitaryTime(mergedCalanders[len(mergedCalanders)-1].End)) {
				mergedCalanders[len(mergedCalanders)-1].End = max(mergedCalanders[len(mergedCalanders)-1].End, b[r].End)
			} else {
				mergedCalanders = append(mergedCalanders, b[r])
			}
			r++
		}
	}

	for l < len(a) {
		if a[l].Start < mergedCalanders[len(mergedCalanders)-1].End {
			mergedCalanders[len(mergedCalanders)-1].End = max(mergedCalanders[len(mergedCalanders)-1].End, a[l].End)
		} else {
			mergedCalanders = append(mergedCalanders, a[l])
		}
		l++
	}

	for r < len(b) {
		if b[r].Start < mergedCalanders[len(mergedCalanders)-1].End {
			mergedCalanders[len(mergedCalanders)-1].End = max(mergedCalanders[len(mergedCalanders)-1].End, b[r].End)

		} else {
			mergedCalanders = append(mergedCalanders, b[r])

		}
	}

	return mergedCalanders
}

type MilitaryTime string

func (m MilitaryTime) ToMinutes() int {
	splits := strings.Split(string(m), ":")
	hours, minutes := splits[0], splits[1]

	if strings.HasPrefix(hours, "0") {
		hours = string(hours[1:])
	}
	if strings.HasPrefix(minutes, "0") {
		minutes = string(minutes[1:])
	}

	hoursNum, _ := strconv.Atoi(hours)
	minutesNum, _ := strconv.Atoi(minutes)

	return (hoursNum * 60) + minutesNum
}

func (m MilitaryTime) AddMinutes(minutes int) MilitaryTime {
	mm := m.ToMinutes()
	mm = (mm + minutes) % (24 * 60)
	return minutesToMilitaryTime(mm)
}

func (m MilitaryTime) Add(other MilitaryTime) MilitaryTime {
	mm, mo := m.ToMinutes(), other.ToMinutes()
	mm += mo
	return minutesToMilitaryTime(mm)
}

func (m MilitaryTime) LessThan(other MilitaryTime) bool {
	mm, mo := m.ToMinutes(), other.ToMinutes()
	return mm < mo
}

func (m MilitaryTime) Equals(other MilitaryTime) bool {
	mm, mo := m.ToMinutes(), other.ToMinutes()
	return mm == mo
}

func minutesToMilitaryTime(minutes int) MilitaryTime {
	minutes = minutes % (24 * 60)
	hours := minutes / 60
	minutes = minutes % 60

	hourStr := strconv.Itoa(hours)
	minutesStr := strconv.Itoa(minutes)

	if len(hourStr) == 1 {
		hourStr = fmt.Sprintf("0%s", hourStr)
	}

	if len(minutesStr) == 1 {
		minutesStr = fmt.Sprintf("0%s", minutesStr)
	}

	return MilitaryTime(fmt.Sprintf("%s:%s", hourStr, minutesStr))
}

func max(a, b string) string {
	am, bm := MilitaryTime(a).ToMinutes(), MilitaryTime(b).ToMinutes()
	if am > bm {
		return a
	}
	return b
}
