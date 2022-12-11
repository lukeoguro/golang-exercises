// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
package kata

func isEven(num int) bool {
	return num%2 == 0
}

func FindOutlier(integers []int) int {
	even := 0
	for _, num := range integers[:3] {
		if isEven(num) {
			even += 1
		}
	}

	var outlier int
	for _, num := range integers {
		if (even <= 1 && isEven(num)) || (even > 1 && !isEven(num)) {
			outlier = num
		}
	}

	return outlier
}
