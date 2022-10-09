package binary

import (
	"math"
)

func Rotationv2(haystack []int) uint {
	min := uint(0)
	max := uint(len(haystack))
	for min < max {
		m := uint(math.Floor(float64(min + (max-min)/2)))
		if uint(haystack[uint(math.Mod(float64(m-1), float64(max)))]) > uint(haystack[m]) {
			return m
		}

		if uint(haystack[uint(math.Mod(float64(m+1), float64(max)))]) < uint(haystack[m]) {
			return uint(math.Mod(float64(m+1), float64(max)))
		}

		if uint(haystack[min]) < uint(haystack[m]) {
			min = uint(math.Mod(float64(m+1), float64(max)))
		} else {
			max = m
		}
	}

	return 0
}
