// https://www.codewars.com/kata/5679aa472b8f57fb8c000047/train/go
package kata

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func FindEvenIndex(arr []int) int {
	leftSum := 0
	rightSum := sum(arr)

	for i := 0; i < len(arr); i += 1 {
		rightSum -= arr[i]

		if leftSum == rightSum {
			return i
		}

		leftSum += arr[i]
	}

	return -1
}
