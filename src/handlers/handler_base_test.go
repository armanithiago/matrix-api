package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//type requestTestCase struct {
//	name            string
//	fileName        string
//	convertedMatrix [][]int
//	method          string
//	err             error
//}

type matrixConvertTestCase struct {
	name            string
	strMatrix       [][]string
	convertedMatrix [][]int
	err             error
}

func TestGetCsvFileFromRequest(t *testing.T) {
	requestTestCases := []RequestTestCase{
		{"3x3 Matrix, 1 to 9", "POST", "mock/test/getCsvFileFromRequest", true, "../../assets/matrix_3x3.csv", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "", 0, nil},
		{"3x3 Matrix, 0 to -8", "POST", "mock/test/getCsvFileFromRequest", true, "../../assets/matrix_3x3_negative.csv", [][]int{{0, -1, -2}, {-3, -4, -5}, {-6, -7, -8}}, "", 0, nil},
		{"2x3 Matrix, Non-Quadratic", "POST", "mock/test/getCsvFileFromRequest", true, "../../assets/matrix_2x3_non_quadratic.csv", nil, "", 0, errors.New(NOT_QUADRATIC)},
		{"3x3 Matrix, Non-Integer Characters", "POST", "mock/test/getCsvFileFromRequest", true, "../../assets/matrix_3x3_non_integer_characters.csv", nil, "", http.StatusBadRequest, errors.New(INVALID_CHARACTERS)},
		{"3x3 Matrix, Wrong Request Method", "GET", "mock/test/getCsvFileFromRequest", true, "../../assets/matrix_3x3.csv", nil, "", 0, errors.New(NOT_ALLOWED)},
		{"No attachment request", "POST", "mock/test/getCsvFileFromRequest", false, "", nil, "", 0, errors.New(INVALID_INPUT_TYPE)},
	}

	for _, testCase := range requestTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			request, err := BuildRequest(testCase)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			m, err := GetCsvFileFromRequest(rr, request)

			if reflect.DeepEqual(err, testCase.Err) == false {
				t.Fatalf("Got Error %v, Expected %v", err, testCase.Err)
			}
			if reflect.DeepEqual(m, testCase.ConvertedMatrix) == false {
				t.Fatalf("Got Matrix %v, Expected %v", m, testCase.ConvertedMatrix)
			}
		})
	}
}

func TestConvertMatrixToInt(t *testing.T) {
	matrixConvertTestCases := []matrixConvertTestCase{
		{"2x2 Matrix", [][]string{{"1", "2"}, {"3", "4"}}, [][]int{{1, 2}, {3, 4}}, nil},
		{"3x3 Matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, nil},
		{"2x3 Matrix", [][]string{{"1", "2", "3"}, {"4", "5", "6"}}, [][]int{{1, 2, 3}, {4, 5, 6}}, nil},
		{"3x3 Matrix, 0 to -8", [][]string{{"0", "-1", "-2"}, {"-3", "-4", "-5"}, {"-6", "-7", "-8"}}, [][]int{{0, -1, -2}, {-3, -4, -5}, {-6, -7, -8}}, nil},
		{"2x2 Matrix, Non-Integer Characters", [][]string{{"R", "q"}, {"$", "c"}}, nil, errors.New("only integers allowed on input file")},
	}
	for _, testCase := range matrixConvertTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			m, err := convertMatrixToInt(testCase.strMatrix)

			if reflect.DeepEqual(err, testCase.err) == false {
				t.Fatalf("Got Error %v, Expected %v", err, testCase.err)
			}
			if reflect.DeepEqual(m, testCase.convertedMatrix) == false {
				t.Fatalf("Got Matrix %v, Expected %v", m, testCase.convertedMatrix)
			}
		})
	}
}
