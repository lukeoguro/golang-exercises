// https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go
package kata

func CountPositivesSumNegatives(numbers []int) []int {
	sum := 0
	count := 0

	for _, num := range numbers {
		if num > 0 {
			count += 1
		} else if num < 0 {
			sum += num
		}
	}

	return []int{count, sum}
}
