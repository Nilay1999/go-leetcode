package main

import "fmt"

func intToRoman(input int) string {
	numArr := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanArr := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	i := 0
	for input > 0 {
		if input < numArr[i] {
			i += 1
		} else {
			result += romanArr[i]
			input -= numArr[i]
		}

	}
	return result
}

func main() {
	input := 1994
	fmt.Print(intToRoman(input))
}
