// https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go
package kata

import "math"

func FindNextSquare(sq int64) int64 {
	sqrt := math.Sqrt(float64(sq))
	if math.Trunc(sqrt) != sqrt {
		return -1
	}

	return int64(math.Pow(sqrt+1, 2))
}
