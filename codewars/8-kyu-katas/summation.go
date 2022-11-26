// https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/go
package kata

func Summation(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
