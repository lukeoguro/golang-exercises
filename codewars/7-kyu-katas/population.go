// https://www.codewars.com/kata/563b662a59afc2b5120000c6/train/go
package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	population := p0
	year := 0

	for population < p {
		population += int(float64(population) * percent / 100)
		population += aug
		year += 1
	}

	return year
}
