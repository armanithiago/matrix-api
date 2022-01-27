package handlers

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type requestTestCase struct {
	name            string
	fileName        string
	convertedMatrix [][]int
	method          string
	err             error
}

type matrixConvertTestCase struct {
	name            string
	strMatrix       [][]string
	convertedMatrix [][]int
	err             error
}

func TestGetCsvFileFromRequest(t *testing.T) {
	requestTestCases := []requestTestCase{
		{"3x3 Matrix, 1 to 9", "../../assets/matrix_3x3.csv", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "POST", nil},
		{"3x3 Matrix, 0 to -8", "../../assets/matrix_3x3_negative.csv", [][]int{{0, -1, -2}, {-3, -4, -5}, {-6, -7, -8}}, "POST", nil},
		{"2x3 Matrix, Non-Quadratic", "../../assets/matrix_2x3_non_quadratic.csv", nil, "POST", errors.New("invalid input. File should have same number of rows and columns")},
		{"3x3 Matrix, Non-Integer Characters", "../../assets/matrix_3x3_non_integer_characters.csv", nil, "POST", errors.New("only integers allowed on input file")},
		{"3x3 Matrix, Wrong Request Method", "../../assets/matrix_3x3.csv", nil, "GET", errors.New("method not allowed")},
	}

	for _, testCase := range requestTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			file, _ := os.Open("../testfiles/" + testCase.fileName)
			defer file.Close()

			var requestBody bytes.Buffer
			multipartWriter := multipart.NewWriter(&requestBody)
			fileWriter, err := multipartWriter.CreateFormFile("file", testCase.fileName)
			if err != nil {
				t.Fatal(err)
			}

			_, err = io.Copy(fileWriter, file)
			if err != nil {
				t.Fatal(err)
			}
			multipartWriter.Close()

			request, err := http.NewRequest(testCase.method, "mock/test/getCsvFileFromRequest", &requestBody)
			request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			m, err := GetCsvFileFromRequest(rr, request)

			if reflect.DeepEqual(err, testCase.err) == false {
				t.Fatalf("Got Error %v, Expected %v", err, testCase.err)
			}
			if reflect.DeepEqual(m, testCase.convertedMatrix) == false {
				t.Fatalf("Got Matrix %v, Expected %v", m, testCase.convertedMatrix)
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
