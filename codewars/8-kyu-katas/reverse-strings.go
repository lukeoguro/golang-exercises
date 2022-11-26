// https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
package kata

func Solution(word string) (reversed string) {
	for _, char := range word {
		reversed = string(char) + reversed
	}
	return
}
