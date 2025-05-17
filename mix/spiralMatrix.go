package mix

func SpiralMatrix(input [][]int) []int {
	res := []int{}
	top, bottom, left, right := 0, len(input)-1, 0, len(input[0])-1
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			res = append(res, input[top][j])
		}
		top += 1
		for j := top; j <= bottom; j++ {
			res = append(res, input[j][right])
		}
		right -= 1

		if !(left < right+1 && top < bottom+1) {
			break
		}

		for j := right; j > left-1; j-- {
			res = append(res, input[bottom][j])
		}
		bottom -= 1
		for j := bottom; j > top-1; j-- {
			res = append(res, input[j][left])
		}
		left += 1
	}

	return res
}
