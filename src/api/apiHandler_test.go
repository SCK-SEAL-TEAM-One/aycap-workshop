package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_StartDay_4_StartMonth_1_StartYear_2018_EndDay_4_EndMonth_1_EndYear_2018_Should_Be_DurationResponse(t *testing.T) {
	url := "/CalculationDurationBetweenTwoDate?start_day=20&start_month=8&start_year=1986&end_day=21&end_month=6&end_year=2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"days":"15,082 days"}`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}
}
