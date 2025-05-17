package mix

func RotateImage(arr [][]int) [][]int {
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
