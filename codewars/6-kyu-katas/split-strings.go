// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
package kata

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str += "_"
	}

	pairs := make([]string, len(str)/2)
	for i := 0; i < len(pairs); i += 1 {
		pairs[i] = str[i*2 : i*2+2]
	}

	return pairs
}
