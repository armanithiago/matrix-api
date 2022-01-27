package format

import (
	"errors"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	requestTestCases := []handlers.RequestTestCase{
		{"3x3 Matrix, 1 to 9", "POST", "localhost:8080/echo", true, "../../../assets/matrix_3x3.csv", nil, "1,2,3\n4,5,6\n7,8,9", http.StatusOK, nil},
		{"3x3 Matrix, 0 to -8", "POST", "localhost:8080/echo", true, "../../../assets/matrix_3x3_negative.csv", nil, "0,-1,-2\n-3,-4,-5\n-6,-7,-8", http.StatusOK, nil},
		{"2x3 Matrix, Non-Quadratic", "POST", "localhost:8080/echo", true, "../../../assets/matrix_2x3_non_quadratic.csv", nil, "Error: " + handlers.NOT_QUADRATIC, http.StatusBadRequest, errors.New(handlers.NOT_QUADRATIC)},
		{"3x3 Matrix, Non-Integer Characters", "POST", "localhost:8080/echo", true, "../../../assets/matrix_3x3_non_integer_characters.csv", nil, "Error: " + handlers.INVALID_CHARACTERS, http.StatusBadRequest, errors.New(handlers.INVALID_CHARACTERS)},
		{"3x3 Matrix, Wrong Request Method", "GET", "localhost:8080/echo", true, "../../../assets/matrix_3x3.csv", nil, "Error: " + handlers.NOT_ALLOWED, http.StatusBadRequest, errors.New(handlers.NOT_ALLOWED)},
		{"No attachment request", "POST", "localhost:8080/echo", false, "", nil, "Error: " + handlers.INVALID_INPUT_TYPE, http.StatusBadRequest, errors.New(handlers.INVALID_INPUT_TYPE)},
	}

	for _, testCase := range requestTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			request, err := handlers.BuildRequest(testCase)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Echo)
			handler.ServeHTTP(rr, request)

			if status := rr.Code; status != testCase.ExpectedStatus {
				t.Errorf("Handler returned wrong status code: got %v expected %v",
					status, testCase.ExpectedStatus)
			}

			if rr.Body.String() != testCase.ExpectedResult {
				t.Errorf("Handler returned unexpected body: got \n%v\n want \n%v\n",
					rr.Body.String(), testCase.ExpectedResult)
			}
		})
	}
}
