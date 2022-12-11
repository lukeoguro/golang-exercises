// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go
package kata

func FindUniq(arr []float32) float32 {
	lookup := map[float32]int{}

	for _, v := range arr {
		lookup[v] += 1
	}

	var uniq float32
	for k, v := range lookup {
		if v == 1 {
			uniq = k
		}
	}

	return uniq
}
