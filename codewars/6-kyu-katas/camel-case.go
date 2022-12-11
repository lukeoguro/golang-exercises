// https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
package kata

import "strings"

func ToCamelCase(s string) string {
	newStr := ""

	afterDelimiter := false
	for _, rn := range s {
		char := string(rn)
		if char == "-" || char == "_" {
			afterDelimiter = true
		} else if afterDelimiter {
			newStr += strings.ToUpper(char)
			afterDelimiter = false
		} else {
			newStr += char
		}
	}

	return newStr
}
