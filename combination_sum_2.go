package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	var recursion func(idx int, target int, temp []int)

	recursion = func(idx int, target int, temp []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, temp...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			recursion(i+1, target-candidates[i], temp)
			temp = temp[:len(temp)-1]
		}
	}
	recursion(0, target, []int{})
	return ans
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Print(combinationSum2(candidates, target))
}
