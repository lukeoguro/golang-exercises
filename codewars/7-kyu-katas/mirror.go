// https://www.codewars.com/kata/56dbe0e313c2f63be4000b25/train/go
package kata

import (
	"strings"
)

func VertMirror(s string) string {
	substrs := strings.Split(s, "\n")
	mirroredSubstrs := make([]string, len(substrs))

	for i, substr := range substrs {
		mirroredSubstrs[i] = reverseStr(substr)
	}

	return strings.Join(mirroredSubstrs, "\n")
}

func HorMirror(s string) string {
	substrs := strings.Split(s, "\n")
	reverseSlice(substrs)
	return strings.Join(substrs, "\n")
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func reverseStr(str string) (reversedStr string) {
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}

	return
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}
