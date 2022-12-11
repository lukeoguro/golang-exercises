// https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec/train/go
package kata

import "strconv"

func Persistence(n int) int {
	times := 0
	str := strconv.Itoa(n)
	for len(str) > 1 {
		persistence := 1
		for _, rn := range str {
			num, _ := strconv.Atoi(string(rn))
			persistence *= num
		}

		str = strconv.Itoa(persistence)
		times += 1
	}

	return times
}
