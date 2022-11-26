// https://www.codewars.com/kata/59c5f4e9d751df43cf000035/train/go
package kata

import "regexp"

func Solve(s string) int {
	vowels, _ := regexp.Compile("[aeiou]+")
	matches := vowels.FindAllString(s, -1)

	length := 0
	for _, match := range matches {
		if len(match) > length {
			length = len(match)
		}
	}

	return length
}
