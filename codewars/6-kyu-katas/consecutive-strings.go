// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go
package kata

func LongestConsec(strArr []string, k int) string {
	n := len(strArr)
	if n == 0 || k > n || k <= 0 {
		return ""
	}

	longest := ""
	for i := 0; i <= n-k; i += 1 {
		str := ""
		for j := 0; j < k; j += 1 {
			str += strArr[i+j]
		}

		if len(str) > len(longest) {
			longest = str
		}
	}

	return longest
}
