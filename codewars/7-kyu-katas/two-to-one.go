// https://www.codewars.com/kata/5656b6906de340bd1b0000ac/train/go
package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	set := map[string]bool{}
	for _, char := range s1 {
		set[string(char)] = true
	}
	for _, char := range s2 {
		set[string(char)] = true
	}

	uniqChars := make([]string, len(set))
	i := 0
	for k, _ := range set {
		uniqChars[i] = k
		i++
	}

	sort.Strings(uniqChars)
	return strings.Join(uniqChars, "")
}
