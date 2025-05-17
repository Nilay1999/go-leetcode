package main

import (
	"math"
)

func WaterContainer(arr []float64) float64 {
	max := 0.0
	l, r := 0, len(arr)-1
	for l < r {
		total := math.Min(arr[l], arr[r]) * float64(r-l)
		max = math.Max(total, max)
		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
