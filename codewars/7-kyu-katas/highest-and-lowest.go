// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go
package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numStrs := strings.Split(in, " ")

	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}

	min, max := minMax(nums)
	return fmt.Sprintf("%v %v", max, min)
}

func minMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return min, max
}
