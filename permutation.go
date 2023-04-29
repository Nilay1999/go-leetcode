package main

import "fmt"

func contains(nums []int, element int) bool {
	for _, val := range nums {
		if val == element {
			return true
		}
	}
	return false
}

func permute(nums []int) [][]int {
	ans := [][]int{}

	var recursion func(temp []int)
	recursion = func(temp []int) {
		if len(temp) == len(nums) {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if contains(temp, nums[i]) {
				continue
			}
			temp = append(temp, nums[i])
			recursion(temp)
			temp = temp[:len(temp)-1]
		}
	}
	recursion([]int{})
	return ans
}

func main() {
	input := []int{1, 2, 3}
	fmt.Print(permute(input))
}
