package db

import (
	"fmt"
	"regexp"
	"strings"
)

//Credit: https://gist.github.com/adharris/4163702
var (
	// unquoted array values must not contain: (" , \ { } whitespace NULL)
	// and must be at least one char
	unquotedChar  = `[^",\\{}\s]`
	unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)

	// quoted array values are surrounded by double quotes, can be any
	// character except " or \, which must be backslash escaped:
	quotedChar  = `[^"\\]|\\"|\\\\`
	quotedValue = fmt.Sprintf("\"(%s)*\"", quotedChar)

	// an array value may be either quoted or unquoted:
	arrayValue = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)

	// Array values are separated with a comma IF there is more than one value:
	arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))

	valueIndex int
)

// ParseArray is a function that will allow you to extract an array of strings for a Postgress Array Type
func ParseArray(array string) []string {
	var results []string
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		// the string _might_ be wrapped in quotes, so trim them:
		s = strings.Trim(s, "\",")
		results = append(results, s)
	}
	return results

}

// CreateStringArray is a function that will create a string formatted to be used for Postgres Array types
func CreateStringArray(array []string) string {
	if len(array) == 0 {
		return "{}"
	}
	results := []string{}
	for _, v := range array {
		results = append(results, strings.Replace(v, "\"", "\\\"", -1))
	}
	return "{\"" + strings.Join(results, "\",\"") + "\"}"
}
