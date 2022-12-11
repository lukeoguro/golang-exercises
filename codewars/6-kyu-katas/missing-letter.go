// https://www.codewars.com/kata/5839edaa6754d6fec10000a2/train/go
package kata

func FindMissingLetter(chars []rune) rune {
	var missing rune
	for i := 0; i < len(chars)-1; i += 1 {
		if chars[i]+1 != chars[i+1] {
			missing = chars[i] + 1
		}
	}

	return missing
}
