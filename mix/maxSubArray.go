package mix

func MaxSubArray(arr []int) int {
	total, max := 0, -1000000
	for i := 0; i < len(arr); i++ {
		total += arr[i]
		if max < total {
			max = total
		}
		if total < 0 {
			total = 0
		}
	}
	return max
}
