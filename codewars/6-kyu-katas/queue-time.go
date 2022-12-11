// https://www.codewars.com/kata/57b06f90e298a7b53d000a86/train/go
package kata

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func QueueTime(queue []int, n int) int {
	tills := make([]int, n)
	time := 0

	for len(queue) != 0 || sum(tills) != 0 {
		for i, v := range tills {
			if v == 0 && len(queue) > 0 {
				tills[i] = queue[0]
				queue = queue[1:]
			}

			if tills[i] > 0 {
				tills[i] -= 1
			}

		}
		time += 1
	}

	return time
}
