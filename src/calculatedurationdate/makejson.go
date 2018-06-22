package calculatedurationdate

import "fmt"

type Duration struct {
	From    string `json:"form"`
	To      string `json:"to"`
	Days    string `json:"days"`
	Seconds string `json:"seconds"`
	Minutes string `json:"minutes"`
	Hours   string `json:"hours"`
}

func MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear int) Duration {
	//totalDay := calculeteTotalDay(startDay, startMonth, startYear, endDay, endMonth, endYear)
	hours := TransformDaysToHour(15082)
	//minutes := TransformDayToMinutes(hounrs)
	seconds := TransformMinutesToSeconds(21718080)
	secondsResult := fmt.Sprintf("%s %s", AddComma(int64(seconds)), "Seconds")

	return Duration{
		//from:    TransformDayToFullDate(startDay, startMonth, startYear),
		//	to:      TransformDayFullDate(endDay, endMonth, endYear),
		//days:    fmt.Sprintf(" %s %s", AddComma(totalDay), "Days"),
		Seconds: secondsResult,
		//minutes: fmt.Sprintf(" %s %s", AddComma(minutes), "Minutes"),
		Hours: fmt.Sprintf("%s %s", AddComma(hours), "Hours"),
	}
}
