package calculatedurationdate

import (
	"bytes"
	"strconv"
)

func ReverseNumber(number string) string {
	reverse := []rune(number)
	for i, j := 0, len(reverse)-1; i < len(reverse)/2; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}

func AddComma(number int64) string {
	numberString := strconv.FormatInt(number, 10)
	var buffer bytes.Buffer
	reverseNumbers := ReverseNumber(numberString)
	count := 0
	for i := 0; i < len(numberString); i++ {
		count++
		buffer.WriteString(reverseNumbers[i : i+1])

		if count == 3 && i != len(numberString)-1 {
			buffer.WriteString(",")
			count = 0
		}
	}
	return ReverseNumber(buffer.String())
}
