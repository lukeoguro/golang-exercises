// https://www.codewars.com/kata/55c45be3b2079eccff00010f/train/go
package kata

import (
	"regexp"
	"sort"
	"strings"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")

	re, _ := regexp.Compile("[1-9]")
	sort.Slice(words, func(i, j int) bool {
		iNum := re.FindString(words[i])
		jNum := re.FindString(words[j])

		return iNum < jNum
	})

	return strings.Join(words, " ")
}
