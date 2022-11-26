// https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go
package kata

func PositiveSum(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return
}
