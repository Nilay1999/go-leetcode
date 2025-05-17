package main

func CombinationSum(candidates []int, target int) [][]int {
	ans, idx, temp := [][]int{}, 0, []int{}
	var recursion func(idx int, target int, temp []int)
	recursion = func(idx int, target int, temp []int) {
		if idx == len(candidates) {
			if target == 0 {
				ans = append(ans, append([]int{}, temp...))
			}
			return
		}
		if candidates[idx] <= target {
			temp = append(temp, candidates[idx])
			recursion(idx, target-candidates[idx], temp)
			temp = temp[:len(temp)-1]
		}
		recursion(idx+1, target, temp)

	}
	recursion(idx, target, temp)
	return ans
}
