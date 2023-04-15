package main

import (
	"fmt"
	"math"
)

func searchInRotatedArray(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := int(math.Floor(float64((l + r) / 2)))
	for l <= r {
		if nums[mid] == target {
			return mid
		} else if nums[l] == target {
			return l
		} else if nums[r] == target {
			return r
		}
		if nums[l] < nums[mid] {
			if nums[l] < target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func main() {

	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	fmt.Print(searchInRotatedArray(input, target))
}
