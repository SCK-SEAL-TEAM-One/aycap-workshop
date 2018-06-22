package calculatedurationdate

import (
	"fmt"
	"time"
)

func CalculateTotalDayResult(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)
	fmt.Printf("test startDate %v:", startDate)
	fmt.Printf("test endDate %v:", endDate)
	//fmt.Printf(endDate.sub(startDate))
	difference := endDate.Sub(startDate)
	totalDay := (difference.Hours() / 24) + 1
	//fmt.Printf("test difference %d:", hour)
	return int(totalDay)
}
