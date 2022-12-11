// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
package kata

import (
	"strconv"
)

func DigitalRoot(n int) int {
	sum := 0
	for _, digitStr := range strconv.Itoa(n) {
		digit, _ := strconv.Atoi(string(digitStr))
		sum += digit
	}

	if sum < 10 {
		return sum
	} else {
		return DigitalRoot(sum)
	}
}
