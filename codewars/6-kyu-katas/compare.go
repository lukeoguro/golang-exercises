// https://www.codewars.com/kata/550498447451fbbd7600041c/train/go
package kata

func Comp(arr1 []int, arr2 []int) bool {
	if arr1 == nil || arr2 == nil {
		return false
	}

	if len(arr1) != len(arr2) {
		return false
	}

	lookup := map[int]int{}

	for _, v := range arr2 {
		lookup[v] += 1
	}

	for _, v := range arr1 {
		if lookup[v*v] == 0 {
			return false
		} else {
			lookup[v*v] -= 1
		}
	}

	return true
}
