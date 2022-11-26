// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go
package kata

func GetMiddle(s string) string {
	midIndex := len(s) / 2

	if len(s)%2 == 1 {
		return string(s[midIndex])
	} else {
		return s[midIndex-1 : midIndex+1]
	}
}
