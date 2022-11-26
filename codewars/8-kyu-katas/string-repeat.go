// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e/train/go
package kata

import (
	s "strings"
)

func RepeatStr(repetitions int, value string) string {
	return s.Repeat(value, repetitions)
}
