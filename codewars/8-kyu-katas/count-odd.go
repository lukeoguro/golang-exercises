// https://www.codewars.com/kata/59342039eb450e39970000a6/train/go
package kata

func OddCount(n int) (count int) {
	if n%2 == 0 {
		return n / 2
	} else {
		return (n - 1) / 2
	}
}
