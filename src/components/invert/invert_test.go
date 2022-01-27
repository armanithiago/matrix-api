package invert

import (
	"testing"
)

type testCase struct {
	name     string
	input    [][]int
	expected string
}

func TestInvert(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int{{1, 2}, {3, 4}}, "1,3\n2,4"},
		{"3x3 Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,4,7\n2,5,8\n3,6,9"},
		{"0x0 Matrix", [][]int{}, ""},
		{"1x1 Matrix", [][]int{{1}}, "1"},
		{"3x3 Matrix with negative values", [][]int{{1, -2, 3}, {4, 5, -6}, {-7, 8, 9}}, "1,4,-7\n-2,5,8\n3,-6,9"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Execute(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not inverted properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
