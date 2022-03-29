package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalenderMatching(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		a, b            []StringMeeting
		boundA, boundB  StringMeeting
		meetingDuration int
		expectedSlots   []StringMeeting
	}{
		{
			name: "algoexpert_example",
			a: []StringMeeting{
				{
					Start: "09:00",
					End:   "10:30",
				},
				{
					Start: "12:00",
					End:   "13:00",
				},
				{
					Start: "16:00",
					End:   "18:00",
				},
			},
			boundA: StringMeeting{
				Start: "09:00",
				End:   "20:00",
			},
			b: []StringMeeting{
				{
					Start: "10:00",
					End:   "11:30",
				},
				{
					Start: "12:30",
					End:   "14:30",
				},
				{
					Start: "14:30",
					End:   "15:00",
				},
				{
					Start: "16:00",
					End:   "17:00",
				},
			},
			boundB: StringMeeting{
				Start: "10:00",
				End:   "18:30",
			},
			meetingDuration: 30,
			expectedSlots: []StringMeeting{
				{
					Start: "11:30",
					End:   "12:00",
				},
				{
					Start: "15:00",
					End:   "16:00",
				},
				{
					Start: "18:00",
					End:   "18:30",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := CalendarMatching(tt.a, tt.boundA, tt.b, tt.boundB, tt.meetingDuration)

			assert.Equal(t, tt.expectedSlots, res)
		})
	}
}

func TestMergeCalenders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                    string
		a, b                    []StringMeeting
		expectedMergedCalenders []StringMeeting
	}{
		{
			name:                    "empty_calenders",
			a:                       []StringMeeting{},
			b:                       []StringMeeting{},
			expectedMergedCalenders: []StringMeeting{},
		},
		{
			name: "single_empty_calender",
			a:    []StringMeeting{},
			b: []StringMeeting{
				{
					Start: "9:00",
					End:   "10:00",
				},
			},
			expectedMergedCalenders: []StringMeeting{
				{
					Start: "9:00",
					End:   "10:00",
				},
			},
		},
		{
			name: "neither_empty_with_merge",
			a: []StringMeeting{
				{
					Start: "10:00",
					End:   "11:00",
				},
				{
					Start: "13:00",
					End:   "16:00",
				},
			},
			b: []StringMeeting{
				{
					Start: "9:00",
					End:   "11:00",
				},
				{
					Start: "12:00",
					End:   "12:30",
				},
				{
					Start: "14:00",
					End:   "17:27",
				},
			},
			expectedMergedCalenders: []StringMeeting{
				{
					Start: "9:00",
					End:   "11:00",
				},
				{
					Start: "12:00",
					End:   "12:30",
				},
				{
					Start: "13:00",
					End:   "17:27",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			merged := mergeCalanders(tt.a, tt.b)

			assert.Equal(t, tt.expectedMergedCalenders, merged)
		})
	}
}

func TestMilitaryTime_ToMinutes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "zero_time",
			input:          "00:00",
			expectedOutput: 0,
		},
		{
			name:           "end_time",
			input:          "23:59",
			expectedOutput: 1439,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := MilitaryTime(tt.input).ToMinutes()

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}

func TestMilitaryTime_AddMinutes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		time           string
		minutesToAdd   int
		expectedOutput MilitaryTime
	}{
		{
			name:           "add_no_minutes_to_no_time",
			time:           "00:00",
			minutesToAdd:   0,
			expectedOutput: MilitaryTime("00:00"),
		},
		{
			name:           "add_hour_to_no_time",
			time:           "00:00",
			minutesToAdd:   60,
			expectedOutput: MilitaryTime("01:00"),
		},
		{
			name:           "add_minutes_to_no_time_to_boundary",
			time:           "00:00",
			minutesToAdd:   1439,
			expectedOutput: MilitaryTime("23:59"),
		},
		{
			name:           "add_minutes_to_time_over_boundary",
			time:           "23:30",
			minutesToAdd:   90,
			expectedOutput: MilitaryTime("01:00"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := MilitaryTime(tt.time).AddMinutes(tt.minutesToAdd)

			assert.True(t, res.Equals(tt.expectedOutput))
		})
	}
}

func TestMilitaryTime_LessThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		a, b    string
		isEqual bool
	}{
		{
			name:    "zero_equals_zero",
			a:       "00:00",
			b:       "00:00",
			isEqual: true,
		},
		{
			name:    "non_zero_equals_non_zero",
			a:       "10:21",
			b:       "10:21",
			isEqual: true,
		},
		{
			name:    "minutes_equal_but_not_hour",
			a:       "10:00",
			b:       "11:00",
			isEqual: false,
		},
		{
			name:    "hour_equal_but_not_minutes",
			a:       "11:01",
			b:       "11:00",
			isEqual: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := MilitaryTime(tt.a).Equals(MilitaryTime(tt.b))

			assert.Equal(t, tt.isEqual, res)
		})
	}
}
