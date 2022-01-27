package flatten

import (
	"testing"
)

type testCase struct {
	name     string
	input    [][]int64
	expected string
}

func TestFlatten(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int64{{1, 2}, {3, 4}}, "1,2,3,4"},
		{"3x3 Matrix", [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,2,3,4,5,6,7,8,9"},
		{"0x0 Matrix", [][]int64{}, ""},
		{"1x1 Matrix", [][]int64{{1}}, "1"},
		{"3x3 Matrix with negative values", [][]int64{{1, -2, 3}, {4, 5, -6}, {-7, 8, 9}}, "1,-2,3,4,5,-6,-7,8,9"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Execute(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not flattened properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
