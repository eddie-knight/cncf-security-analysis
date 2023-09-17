package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var REQUEST_COUNTER = 0

func getHttpResponse(url string) *http.Response {
	REQUEST_COUNTER++
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
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
	return days
}

func writeToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(data)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// convertStringToList will convert a string such as {value1, value2, value3} to a list of strings
func convertStringToList(s string) []string {
	var list []string
	// remove the curly braces
	s = s[1 : len(s)-1]
	// split the string by comma
	split := strings.Split(s, ",")
	// trim the spaces
	for _, v := range split {
		list = append(list, strings.TrimSpace(v))
	}
	return list
}

// convertStringToBool will convert a string of "t" or "f" to a boolean
func convertStringToBool(s string) bool {
	return s == "t"
}

// ***
// If CSV needs fixed without making all the API calls
// ***

type CSVContent struct {
	Header []string
	Rows   [][]string
}

func (c *CSVContent) ReadCSV(filename string) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	c.Header = lines[0]
	c.Rows = lines[1:]
}

// change all "Slam22 Participant" column values to true or false depending on whether they contain "yes" or "no" anywhere in the string
func (c *CSVContent) FixBooleanValues() {
	for i, row := range c.Rows {
		if strings.Contains(row[7], "Yes") {
			c.Rows[i][7] = "true"
		} else if strings.Contains(row[7], "No") {
			c.Rows[i][7] = "false"
		}
		log.Printf(c.Rows[i][3] + ": " + c.Rows[i][7])
	}
	// write the contents of the fixed data to a new CSV file
	file, err := os.Create("cloudevents-security-scores.csv")
	file.WriteString(strings.Join(c.Header, ",") + "\n")
	for _, row := range c.Rows {
		file.WriteString(strings.Join(row, ",") + "\n")
	}
	if err != nil {
		log.Fatal(err)
	}

}
