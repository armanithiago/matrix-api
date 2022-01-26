package handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// GetCsvFileFromRequest receive a request and get the file csv attach to it
func GetCsvFileFromRequest(w http.ResponseWriter, r *http.Request) ([][]int, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New(NOT_ALLOWED)
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, errors.New(INVALID_INPUT_TYPE)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, err
	}
	if len(records) <= 1 {
		return nil, errors.New(INVALID_SIZE)
	}

	if len(records) != len(records[0]) {
		return nil, errors.New("NOT_QUADRATIC")
	}

	intMatrix, err := convertMatrixToInt(records)
	if err != nil {
		return nil, err
	}

	return intMatrix, nil
}

//	private convert method responsible to get a string matrix and convert it to int matrix
func convertMatrixToInt(m [][]string) ([][]int, error) {
	var converted [][]int

	for _, row := range m {
		var intRow []int
		for _, column := range row {
			num, err := strconv.Atoi(column)
			if err != nil {
				return nil, errors.New(INVALID_CHARACTERS)
			}
			intRow = append(intRow, num)
		}
		converted = append(converted, intRow)
	}

	if len(converted) != len(converted[0]) {
		return nil, errors.New(NOT_QUADRATIC)
	}

	if len(converted) == 0 {
		return nil, errors.New(INVALID_SIZE)
	}

	return converted, nil
}
