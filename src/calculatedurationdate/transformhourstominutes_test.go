package calculatedurationdate

import "testing"

func Test_TransformHoursToMinutes_Input_361968_Should_Be_21718080(t *testing.T) {
	var totalHours int64 = 361968
	var expected int64 = 21718080

	actual := TransformHoursToMinutes(totalHours)

	if expected != actual {
		t.Errorf("should be equal %d but the actual is %d", expected, actual)
	}
}
