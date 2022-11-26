// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go
package kata

import (
	"strings"
)

func Accum(s string) string {
	s = strings.ToLower(s)
	substrs := make([]string, len(s))

	for i, char := range s {
		substr := strings.Repeat(string(char), i+1)
		substrs[i] = strings.ToUpper(string(substr[0])) + string(substr[1:])
	}

	return strings.Join(substrs, "-")
}
