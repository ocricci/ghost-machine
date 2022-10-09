package binary

import (
	"math"
)

func Rotationv1(haystack []int) uint {
	min := uint(0)
	max := uint(len(haystack))
	needle := uint(haystack[len(haystack)-1] + 1)
	for min < max {
		m := uint(math.Floor(float64(min + (max-min)/2)))
		if uint(haystack[m]) < needle {
			needle = uint(haystack[m])
			max = m
		} else {
			min = m + 1
		}
	}

	return max
}
