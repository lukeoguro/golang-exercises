// https://www.codewars.com/kata/5b180e9fedaa564a7000009a/train/go
package kata

import (
	"regexp"
	"strings"
)

func solve(str string) string {
	lower, _ := regexp.Compile("[a-z]")
	upper, _ := regexp.Compile("[A-Z]")
	lowerCount := len(lower.FindAllString(str, -1))
	upperCount := len(upper.FindAllString(str, -1))

	if lowerCount >= upperCount {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}
}
