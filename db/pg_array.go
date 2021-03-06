package db

import (
	"database/sql/driver"
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
func ParseArray(array string) PGArray {
	var results PGArray
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
func CreateStringArray(array PGArray) string {
	if len(array) == 0 {
		return "{}"
	}
	results := []string{}
	for _, v := range array {
		results = append(results, strings.Replace(v, "\"", "\\\"", -1))
	}
	return "{\"" + strings.Join(results, "\",\"") + "\"}"
}

// PGArray is a type of string slice that is directly usable in PostgreSQL
// queries.
type PGArray []string

// String returns the string slice as a string-representation of a Postgres
// array.
func (a PGArray) String() string {
	return CreateStringArray([]string(a))
}

// Value implements database/sql/driver.Valuer.
func (a PGArray) Value() (driver.Value, error) {
	return a.String(), nil
}

// Scan implements database/sql.Scanner.
func (a *PGArray) Scan(src interface{}) error {
	var str string
	switch t := src.(type) {
	case string:
		str = t
	case []byte:
		str = string(t)
	case nil:
		return nil
	default:
		return fmt.Errorf("null: cannot scan type %T into db.PGArray: %v", src, src)
	}
	*a = ParseArray(str)
	return nil
}
