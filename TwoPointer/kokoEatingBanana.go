package main

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func MinEatingSpeed(piles []int, h int) int {
	l := 1
	r := findMax(piles)
	var res int
	for l < r {
		total := 0
		for i := 0; i < len(piles); i++ {
			var count int
			if piles[i] > l {
				count = int(float64(piles[i] / l))
			} else {
				count = l
			}
			total += int(count)
		}
		if total > h {
			l += 1
		} else {
			r -= 1
		}
	}

	return res
}
