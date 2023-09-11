package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getHttpResponse(url string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	return res
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// get the last day of the month
func getLastDayOfMonth(year int, month int) int {
	dateTime := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
	return dateTime
}

var days []string

// get the last day of the month for every month in the past year
func getLastDaysOfMonths() []string {
	if len(days) > 0 {
		return days
	}
	now := time.Now()
	year := now.Year()
	month := int(now.Month() - 1)
	for i := 0; i < 11; i++ {
		days = append(days, fmt.Sprintf("%d-%02d-%d", year, month, getLastDayOfMonth(year, month)))
		month--
		if month == 0 {
			month = 12
			year--
		}
	}
	fmt.Println(days)
	return days
}
