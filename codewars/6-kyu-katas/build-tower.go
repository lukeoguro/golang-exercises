// https://www.codewars.com/kata/576757b1df89ecf5bd00073b/train/go
package kata

import "strings"

func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)

	spaceCount := nFloors - 1
	starCount := 1

	for i := range tower {
		spaces := strings.Repeat(" ", spaceCount)
		stars := strings.Repeat("*", starCount)
		tower[i] = spaces + stars + spaces

		spaceCount -= 1
		starCount += 2
	}

	return tower
}
