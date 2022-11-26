// https://www.codewars.com/kata/57a049e253ba33ac5e000212/train/go
package kata

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
