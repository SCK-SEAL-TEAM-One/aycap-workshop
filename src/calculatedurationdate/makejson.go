package calculatedurationdate

import "fmt"

type Duration struct {
	from    string `json:"form"`
	to      string `json:"to"`
	days    string `json:"days"`
	seconds string `json:"seconds"`
	minutes string `json:"minutes"`
	hours   string `json:"hours"`
}

func MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear int) Duration {
	//totalDay := calculeteTotalDay(startDay, startMonth, startYear, endDay, endMonth, endYear)
	hours := TransformDaysToHour(15082)
	//minutes := TransformDayToMinutes(hounrs)
	seconds := TransformMinutesToSeconds(21718080)

	return Duration{
		//from:    TransformDayToFullDate(startDay, startMonth, startYear),
		//	to:      TransformDayFullDate(endDay, endMonth, endYear),
		//days:    fmt.Sprintf(" %s %s", AddComma(totalDay), "Days"),
		seconds: fmt.Sprintf("%s %s", AddComma(int64(seconds)), "Seconds"),
		//minutes: fmt.Sprintf(" %s %s", AddComma(minutes), "Minutes"),
		hours: fmt.Sprintf("%s %s", AddComma(hours), "Hours"),
	}
}
