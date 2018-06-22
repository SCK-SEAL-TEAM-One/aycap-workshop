package calculatedurationdate

import "testing"

func Test_CalculateTotalDayResult_Input_7_3_1977_and_21_6_2018_Should_Be_15082(t *testing.T) {
	startDay := 7
	startMonth := 3
	startYear := 1977
	endDay := 21
	endMonth := 6
	endYear := 2018

	expectedDays := 15082
	actualDays := CalculateTotalDayResult(startDay, startMonth, startYear, endDay, endMonth, endYear)

	if expectedDays != actualDays {
		t.Errorf("Should be %d but actual days %d", expectedDays, actualDays)

	}

}
