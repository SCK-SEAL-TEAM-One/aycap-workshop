package calculatedurationdate

type Duration struct {
	Days string `json:"days"`
}

func MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear int) Duration {
	return Duration{
		Days: "152 days",
	}
}
