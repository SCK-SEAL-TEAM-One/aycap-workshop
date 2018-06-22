package calculatedurationdate

import (
	"fmt"
	"time"
)

func TransformDateToFullDate(day, month, year int) string {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	weekDay := date.Weekday().String()
	fullMonth := date.Month()
	return fmt.Sprintf("%s, %d %s %d", weekDay, day, fullMonth, year)
}
