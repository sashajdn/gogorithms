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

// CalendarMatching ...
//
// T -> O(m + n)
// S -> O(m + n)
func CalendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	mergedIntervals := mergeCalanders(calendar1, calendar2)

	var (
		availableSlots = []StringMeeting{}
		startBound     = MilitaryTime(maxInterval(dailyBounds1.Start, dailyBounds2.Start))
		endBound       = MilitaryTime(minInterval(dailyBounds1.End, dailyBounds2.End))
	)
	for i := 0; i < len(mergedIntervals); i++ {
		currentBound := mergedIntervals[i]

		ctm := MilitaryTime(currentBound.Start)
		if !ctm.LessThan(endBound) {
			return availableSlots
		}

		if startBound.LessThan(ctm) {
			if ctm.ToMinutes()-MilitaryTime(startBound).ToMinutes() >= meetingDuration {
				availableSlots = append(availableSlots, StringMeeting{
					Start: string(startBound),
					End:   string(ctm),
				})
			}
		}

		startBound = MilitaryTime(currentBound.End)
	}

	if startBound.LessThan(endBound) {
		availableSlots = append(availableSlots, StringMeeting{
			Start: string(startBound),
			End:   string(endBound),
		})
	}

	return availableSlots
}

func mergeCalanders(a, b []StringMeeting) []StringMeeting {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	var (
		l, r            = 0, 0
		mergedCalanders []StringMeeting
	)
	switch {
	case !MilitaryTime(a[l].Start).LessThan(MilitaryTime(b[r].Start)):
		mergedCalanders = append(mergedCalanders, b[r])
		r++
	default:
		mergedCalanders = append(mergedCalanders, a[l])
		l++
	}

	for l < len(a) && r < len(b) {
		switch {
		case !MilitaryTime(a[l].Start).LessThan(MilitaryTime(b[r].Start)):
			current := b[r]
			r++

			if !MilitaryTime(current.Start).LessThan(MilitaryTime(mergedCalanders[len(mergedCalanders)-1].End)) {
				mergedCalanders = append(mergedCalanders, current)
				continue
			}

			mergedCalanders[len(mergedCalanders)-1].End = maxInterval(mergedCalanders[len(mergedCalanders)-1].End, current.End)
		default:
			current := a[l]
			l++

			if !MilitaryTime(current.Start).LessThan(MilitaryTime(mergedCalanders[len(mergedCalanders)-1].End)) {
				mergedCalanders = append(mergedCalanders, current)
				continue
			}

			mergedCalanders[len(mergedCalanders)-1].End = maxInterval(mergedCalanders[len(mergedCalanders)-1].End, current.End)
		}
	}

	for l < len(a) {
		current := a[l]
		l++

		if current.Start < mergedCalanders[len(mergedCalanders)-1].End {
			mergedCalanders[len(mergedCalanders)-1].End = maxInterval(mergedCalanders[len(mergedCalanders)-1].End, current.End)
			continue
		}
		mergedCalanders = append(mergedCalanders, a[l])
	}

	for r < len(b) {
		current := b[r]
		r++

		if current.Start < mergedCalanders[len(mergedCalanders)-1].End {
			mergedCalanders[len(mergedCalanders)-1].End = maxInterval(mergedCalanders[len(mergedCalanders)-1].End, current.End)
			continue
		}
		mergedCalanders = append(mergedCalanders, b[r])
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

func (m MilitaryTime) Sub(other MilitaryTime) MilitaryTime {
	mm, mo := m.ToMinutes(), other.ToMinutes()
	mm -= mo
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

func maxInterval(a, b string) string {
	am, bm := MilitaryTime(a).ToMinutes(), MilitaryTime(b).ToMinutes()
	if am > bm {
		return a
	}
	return b
}

func minInterval(a, b string) string {
	am, bm := MilitaryTime(a).ToMinutes(), MilitaryTime(b).ToMinutes()
	if am < bm {
		return a
	}
	return b
}
