// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go
package kata

import "strings"

func DuplicateEncode(word string) string {
	lookup := map[string]int{}

	for _, rn := range word {
		char := strings.ToLower(string(rn))
		lookup[char] += 1
	}

	newStr := ""
	for _, rn := range word {
		if lookup[strings.ToLower(string(rn))] == 1 {
			newStr += "("
		} else {
			newStr += ")"
		}
	}

	return newStr
}
