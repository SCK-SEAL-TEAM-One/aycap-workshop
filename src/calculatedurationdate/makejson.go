package calculatedurationdate

import "fmt"

type Duration struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Days    string `json:"days"`
	Seconds string `json:"seconds"`
	Minutes string `json:"minutes"`
	Hours   string `json:"hours"`
}

func MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear int) Duration {
	totalDay := CalculateTotalDayResult(startDay, startMonth, startYear, endDay, endMonth, endYear)
	hours := TransformDaysToHour(totalDay)
	minutes := TransformHoursToMinutes(hours)
	seconds := TransformMinutesToSeconds(minutes)
	secondsResult := fmt.Sprintf("%s %s", AddComma(int64(seconds)), "seconds")

	return Duration{
		From:    TransformDateToFullDate(startDay, startMonth, startYear),
		To:      TransformDateToFullDate(endDay, endMonth, endYear),
		Days:    fmt.Sprintf("%s %s", AddComma(int64(totalDay)), "days"),
		Seconds: secondsResult,
		Minutes: fmt.Sprintf("%s %s", AddComma(minutes), "minutes"),
		Hours:   fmt.Sprintf("%s %s", AddComma(hours), "hours"),
	}
}
