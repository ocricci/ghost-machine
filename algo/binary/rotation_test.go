package binary

import "testing"

func rotationTestCases() [][]int {
	return [][]int{
		[]int{10, 2, 3, 4, 5},
		[]int{9, 10, 11, 2, 3, 4},
		[]int{2, 3, 4, 5}, // no rotation
		[]int{9, 10, 11, 12, 8},
	}
}

func rotationExpectedResults() []int {
	return []int{1, 3, 0, 4}
}

func TestRotation(t *testing.T) {
	expectedResults := rotationExpectedResults()
	for i, testCase := range rotationTestCases() {
		if expectedResults[i] != int(Rotationv1(testCase)) {
			t.Errorf("v1: %d failed, %d expected ", i, expectedResults[i])
		}

		if expectedResults[i] != int(Rotationv2(testCase)) {
			t.Errorf("v2: %d failed, %d expected ", i, expectedResults[i])
		}
	}
}
