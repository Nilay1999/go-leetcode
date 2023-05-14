package main

import "fmt"

func firstAndLastIndex(res []int) []int {
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

func main() {
	input := []int{5, 4, -1, 7, 8}
	//input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Print(firstAndLastIndex(input))
}
