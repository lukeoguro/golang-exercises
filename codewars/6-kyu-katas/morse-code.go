// https://www.codewars.com/kata/54b724efac3d5402db00065e/train/go
package kata

import "strings"

func DecodeMorse(morseCode string) string {
	morseCode = strings.Trim(morseCode, " ")
	words := strings.Split(morseCode, "   ")

	for i, word := range words {
		codes := strings.Split(word, " ")
		var decoded string

		for _, code := range codes {
			decoded += MORSE_CODE[code]
		}

		words[i] = decoded

	}

	return strings.Join(words, " ")
}
