package mix

import (
	"sort"
)

func ThreeSum(arr []int) [][]int {
	total := 0
	sort.Ints(arr)
	res := [][]int{}
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		l, r := i+1, len(arr)-1
		for l < r {
			total = arr[i] + arr[l] + arr[r]
			if total < 0 {
				l += 1
			} else if total > 0 {
				r -= 1
			} else {
				for l < r && arr[l] != arr[i] {
					l += 1
				}
				res = append(res, []int{arr[i], arr[l], arr[r]})
				l += 1
				r -= 1
			}
		}
	}
	return res
}
