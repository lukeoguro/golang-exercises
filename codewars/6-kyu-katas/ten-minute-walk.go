// https://www.codewars.com/kata/54da539698b8a2ad76000228/train/go
package kata

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	long := 0
	lat := 0
	for _, rn := range walk {
		switch rn {
		case 'e':
			lat += 1
		case 'w':
			lat -= 1
		case 'n':
			long += 1
		case 's':
			long -= 1
		}
	}

	return long == 0 && lat == 0
}
