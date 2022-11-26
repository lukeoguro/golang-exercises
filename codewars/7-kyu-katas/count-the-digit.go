// https://www.codewars.com/kata/566fc12495810954b1000030
package kata

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) (count int) {
	dStr := strconv.Itoa(d)

	for i := 0; i <= n; i++ {
		sqStr := strconv.Itoa(i * i)
		count += strings.Count(sqStr, dStr)
	}

	return
}
