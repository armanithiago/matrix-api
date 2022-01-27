package sum

import "testing"

type testCase struct {
	name     string
	input    [][]int64
	expected int64
}

func TestSum(t *testing.T) {
	var testCases = []testCase{
		{"2x2 Matrix", [][]int64{{1, 2}, {3, 4}}, 10},
		{"3x3 Matrix", [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 45},
		{"0x0 Matrix", [][]int64{}, 0},
		{"1x1 Matrix", [][]int64{{1}}, 1},
		{"2x3 Matrix", [][]int64{{1, 2, 3}, {4, 5, 6}}, 21},
		{"3x3 Matrix with negative values", [][]int64{{1, -2, 3}, {4, 5, -6}, {-7, 8, 9}}, 15},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Execute(test.input)

			if result != test.expected {
				t.Errorf("Matrix was not summed properly. Expected \n%v\n and got: \n%v\n", test.expected, result)
			}
		})
	}
}
