// https://www.codewars.com/kata/5277c8a221e209d3f6000b56/train/go
package kata

func ValidBraces(str string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	for _, char := range str {
		if _, exists := pairs[char]; exists {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || char != pairs[stack[len(stack)-1]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
