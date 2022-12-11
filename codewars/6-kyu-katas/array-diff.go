// https://www.codewars.com/kata/523f5d21c841566fde000009/train/go
package kata

func ArrayDiff(a, b []int) []int {
	set := map[int]bool{}
	for _, num := range b {
		set[num] = true
	}

	diffSlice := []int{}
	for _, num := range a {
		if _, exists := set[num]; !exists {
			diffSlice = append(diffSlice, num)
		}
	}

	return diffSlice
}
