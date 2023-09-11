package main

import (
	"strings"
)

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
