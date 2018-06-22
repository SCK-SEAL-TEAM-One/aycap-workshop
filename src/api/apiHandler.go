package api

import (
	"calculatedurationdate"
	"encoding/json"
	"net/http"
	"strconv"
)

func ApiCalculateDate(responseWriter http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	startDay, _ := strconv.Atoi(queryString.Get("start_day"))
	startMonth, _ := strconv.Atoi(queryString.Get("start_month"))
	startYear, _ := strconv.Atoi(queryString.Get("start_year"))
	endDay, _ := strconv.Atoi(queryString.Get("end_day"))
	endMonth, _ := strconv.Atoi(queryString.Get("end_month"))
	endYear, _ := strconv.Atoi(queryString.Get("end_year"))

	durationResponse, _ := json.Marshal(calculatedurationdate.MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear))
	responseWriter.Write(durationResponse)
}
