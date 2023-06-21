package main

import "fmt"

func spiral(n int) [][]int {
	limit := 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	for limit <= n*n {
		for i := left; i <= right; i++ {
			res[top][i] = limit
			limit += 1
		}
		top += 1
		for i := top; i <= bottom; i++ {
			res[i][right] = limit
			limit += 1
		}

		right -= 1
		for i := right; i >= left; i-- {
			res[bottom][i] = limit
			limit += 1
		}

		bottom -= 1
		for i := bottom; i >= top; i-- {
			res[i][left] = limit
			limit += 1
		}
		left += 1
	}

	return res
}

func main() {
	input := 4
	res := spiral(input)
	fmt.Print(res)
}
