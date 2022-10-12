package binary

import (
	"math"
	"testing"
)

func testCases() [][]int {
	return [][]int{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7}}
}

func TestSearchExistingNumber(t *testing.T) {
	for _, testCase := range testCases() {
		m := uint(math.Floor(float64(len(testCase) / 2)))
		result := Search(testCase, testCase[m])
		if !result {
			t.Errorf("search existing number failed")
		}
	}
}
