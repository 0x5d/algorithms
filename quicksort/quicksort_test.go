package sort_test

import (
	"testing"

	"github.com/castillobg/algorithms/quicksort"
)

var testCases = []struct {
	input    []int
	expected []int
}{
	0: {input: []int{9, 3, 6, 1, 0, 5, 4, 0}, expected: []int{0, 0, 1, 3, 4, 5, 6, 9}},
	1: {input: []int{}, expected: []int{}},
	2: {input: []int{234}, expected: []int{234}},
	3: {input: []int{9, 3, 6, 1, 0, 5, 4}, expected: []int{0, 1, 3, 4, 5, 6, 9}},
}

func TestQuicksort(t *testing.T) {
	for i, testCase := range testCases {
		sort.Quicksort(testCase.input)
	inner:
		for j, num := range testCase.input {
			if num != testCase.expected[j] {
				t.Errorf("Test case #%d failed.\nGot: %v\nExpected: %v\n", i, testCase.input, testCase.expected)
				break inner
			}
		}
	}
}
