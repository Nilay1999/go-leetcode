package array

func SecondLargestElement(arr []int) int {
	var max int = -1000

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == max {
			continue
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
