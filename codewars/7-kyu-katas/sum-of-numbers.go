// https://www.codewars.com/kata/55f2b110f61eb01779000053/train/go
package kata

func GetSum(a, b int) int {
	if b < a {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}
