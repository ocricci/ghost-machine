package binary

import (
	"math"
	"testing"
)

func searchTestCases() [][]int {
	return [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{70, 80, 90, 100, 1000, 9000},
	}
}

func TestSearchExistingNumber(t *testing.T) {
	for _, testCase := range searchTestCases() {
		m := uint(math.Floor(float64(len(testCase) / 2)))
		if !Search(testCase, testCase[m]) {
			t.Errorf("Unable to find number in the middle")
		}

		if !Search(testCase, testCase[0]) {
			t.Error("Unable to find the first occurrence")
		}

		if !Search(testCase, testCase[len(testCase) -1]) {
			t.Error("Unable to find the last occurrence")
		}

	}
}

func TestSearchDeadNumber(t *testing.T) {
	for _, testCase := range searchTestCases() {
		if Search(testCase, testCase[len(testCase) - 1] + 1) {
			t.Errorf("Found a dead number !")
		}
	}
}
