package helper

import (
	"log"
	"time"
)

func FormatGoTime(input time.Time) (res time.Time) {
	return input.Truncate(time.Second)
}

func ParsingPgTime(input string) (res time.Time, err error) {
	// postgresTimeStr := "2024-06-28 09:45:55.5777"
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		log.Println("Error parsing time:", err)
		return
	}

	return parsedTime, err
}
