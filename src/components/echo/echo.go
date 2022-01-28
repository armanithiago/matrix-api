package echo

import (
	"fmt"
	"strings"
)

// Execute returns the matrix as a string in matrix format
func Execute(m [][]int64) (s string) {
	var result = ""
	for _, row := range m {
		if result != "" {
			result += "\n"
		}
		strRow := strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ","), "[]")
		result = fmt.Sprintf("%s%s", result, strRow)
	}
	return result
}
