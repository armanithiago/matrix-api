package handlers

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
)

// GetCsvFileFromRequest receive a request and get the file csv attach to it
func GetCsvFileFromRequest(w http.ResponseWriter, r *http.Request) ([][]int64, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New(NOT_ALLOWED)
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, errors.New(INVALID_INPUT_TYPE)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) <= 1 {
		return nil, errors.New(INVALID_SIZE)
	}

	if len(records) != len(records[0]) {
		return nil, errors.New(NOT_QUADRATIC)
	}

	intMatrix, err := convertMatrixToInt(records)
	if err != nil {
		return nil, err
	}

	return intMatrix, nil
}

//	private convert method responsible to get a string matrix and convert it to int matrix
func convertMatrixToInt(m [][]string) ([][]int64, error) {
	var converted [][]int64

	for _, row := range m {
		var intRow []int64
		for _, column := range row {
			num, err := strconv.ParseInt(column, 10, 64)
			if err != nil {
				return nil, errors.New(INVALID_CHARACTERS)
			}
			intRow = append(intRow, num)
		}
		converted = append(converted, intRow)
	}
	return converted, nil
}
