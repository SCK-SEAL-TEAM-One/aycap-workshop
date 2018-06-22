package calculatedurationdate
func TransformMinutesToSeconds(inputMinutes int64)int64{
	const SECONDS_OF_MINUTE = 60
	resultSeconds := inputMinutes*SECONDS_OF_MINUTE
	return int64(resultSeconds)
}
