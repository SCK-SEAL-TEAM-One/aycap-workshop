package calculatedurationdate

import "testing"

func Test_TransformDateToFullDate_input_7_3_1977_should_be_Monday_7_March_1977(t *testing.T) {
	day := 7
	month := 3
	year := 1977
	expectedResult := "Monday, 7 March 1977"
	actualResult := TransformDateToFullDate(day, month, year)
	if actualResult != expectedResult {
		t.Errorf("expected : %s but actual : %s", expectedResult, actualResult)
	}

}
func Test_TransformDateToFullDate_input_21_6_2018_should_be_Thursday_21_June_2018(t *testing.T) {
	day := 21
	month := 6
	year := 2018
	expectedResult := "Thursday, 21 June 2018"
	actualResult := TransformDateToFullDate(day, month, year)
	if actualResult != expectedResult {
		t.Errorf("expected : %s but actual : %s", expectedResult, actualResult)
	}

}
