package echo

import (
	"testing"
)

type testCase struct {
	name     string
	input    [][]int
	expected string
}

func TestEcho(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,2\n3,4"},
		{"3x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,2,3\n4,5,6\n7,8,9"},
		{"0x0 Matrix", [][]int{}, ""},
		{"1x1 Matrix", [][]int{{1}}, "1"},
		{"3x3 Matrix with negative values", [][]int{{1, -2, 3}, {4, 5, -6}, {-7, 8, 9}}, "1,-2,3\n4,5,-6\n-7,8,9"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Execute(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not echoed properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
