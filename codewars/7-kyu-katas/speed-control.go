// https://www.codewars.com/kata/56484848ba95170a8000004d/train/go
package kata

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}

	max := 0
	for i := 1; i < len(x); i++ {
		speed := int((3600 * (x[i] - x[i-1])) / float64(s))
		if speed > max {
			max = speed
		}
	}

	return max
}
