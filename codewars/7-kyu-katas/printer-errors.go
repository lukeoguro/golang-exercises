// https://www.codewars.com/kata/56541980fa08ab47a0000040/train/go
package kata

import (
	"fmt"
	"regexp"
)

func PrinterError(s string) string {
	invalidColor, _ := regexp.Compile("[n-z]")
	matches := invalidColor.FindAllString(s, -1)

	return fmt.Sprintf("%v/%v", len(matches), len(s))
}
