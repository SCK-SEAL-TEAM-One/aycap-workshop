package calculatedurationdate

import "testing"

func Test_TransformDaysToHour_Input_15082_Should_Be_361968(t *testing.T) {
	total_days := 15082
	expected_hours := int64(361968)

	actual_hours := TransformDaysToHour(total_days)

	if expected_hours != actual_hours {
		t.Errorf("expected %d but got it %d", expected_hours, actual_hours)
	}
}
