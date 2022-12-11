// https://www.codewars.com/kata/5264d2b162488dc400000001/train/go
package kata

import "strings"

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}

	return strings.Join(words, " ")
}

func reverse(word string) (reversedWord string) {
	for _, char := range word {
		reversedWord = string(char) + reversedWord
	}

	return
}
