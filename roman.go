package main

import "fmt"

func romanToInt(input string) int {
	mp := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	ans := 0
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && mp[string(input[i])] < mp[string(input[i+1])] {
			ans -= mp[string(input[i])]
		} else {
			ans += mp[string(input[i])]
		}
	}

	return ans
}

func main() {
	input := "LVIII"
	fmt.Print(romanToInt(input))
}
