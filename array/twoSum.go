package main

func TwoSum(array []int, target int) [2]int {
	sub := 0
	mp := make(map[int]int)
	for idx, val := range array {
		sub = target - val
		if res, ok := mp[sub]; ok {
			return [2]int{res, idx}
		}
		mp[val] = idx
	}
	return [2]int{-1, -1}
}
