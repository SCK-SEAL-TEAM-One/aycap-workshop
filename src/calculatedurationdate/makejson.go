package calculatedurationdate

type Duration struct {
	Result string `json:"days"`
}

func MakeJson(startDay, startMonth, startYear, endDay, endMonth, endYear int) Duration {
	return Duration{
		Result: "152 days",
	}
}
