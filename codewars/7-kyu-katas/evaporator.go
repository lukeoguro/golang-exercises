// https://www.codewars.com/kata/5506b230a11c0aeab3000c1f/train/go
package kata

func Evaporator(content float64, evapPerDay int, threshold int) (n int) {
	contentPercentage := float64(1)
	thresholdPercentage := float64(threshold) * 0.01
	multiplier := float64(100-evapPerDay) * 0.01

	for contentPercentage > thresholdPercentage {
		contentPercentage *= multiplier
		n += 1
	}

	return n
}
