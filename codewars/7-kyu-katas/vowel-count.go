// https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
package kata

import "regexp"

func GetCount(str string) int {
	r, _ := regexp.Compile("[aeiou]")
	return len(r.FindAllString(str, -1))
}
