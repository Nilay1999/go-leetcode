package main

import "fmt"

func rotateImage(arr [][]int) [][]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr[0]); j++ {
			temp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = temp
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0])/2; j++ {
			temp := arr[i][j]
			arr[i][j] = arr[i][len(arr[0])-j-1]
			arr[i][len(arr[0])-j-1] = temp
		}
	}
	return arr
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Print(rotateImage(input))

}
