// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9/train/go
package kata

import "strings"

func FindShort(s string) int {
	words := strings.Split(s, " ")

	shortestLen := len(words[0])
	for _, word := range words {
		if len(word) < shortestLen {
			shortestLen = len(word)
		}
	}

	return shortestLen
}
