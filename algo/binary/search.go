package binary

import (
	"math"
)

func Search(haystack []int, needle int) bool {
	min := uint(0)
	max := uint(len(haystack))
	for min < max {
		m := uint(math.Floor(float64(min + (max-min)/2)))
		if haystack[m] == needle {
			return true
		}

		if needle < haystack[m] {
			max = m
		}

		if needle > haystack[m] {
			min = m + 1
		}
	}

	return false
}
