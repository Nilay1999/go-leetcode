package main

import "fmt"

func twoSum(array []int, target int) [2]int {
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

func main() {

	target := 9
	array := []int{2, 7, 11, 15}
	res := twoSum(array[:], target)
	fmt.Print(res)
}
