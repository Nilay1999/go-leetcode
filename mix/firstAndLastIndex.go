package mix

import "fmt"

func FirstAndLastIndex(res []int) []int {
	max, total := -99999, 0
	start, end, glob := 0, 0, 0

	for i := 0; i < len(res); i++ {
		total += res[i]
		if total < 0 {
			total = 0
			start = i + 1
		}
		if max < total {
			glob = start
			max = total
			end = i
		}

	}
	fmt.Print(max)
	return []int{glob, end}
}
