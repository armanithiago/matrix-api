package echo

import (
	"fmt"
	"strings"
)

// Execute returns the print of the matrix exactly as it is
func Execute(m [][]int) (s string) {
	var result string
	for _, row := range m {
		strRow := strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ","), "[]")
		result = fmt.Sprintf("%s%s\n", result, strRow)
	}
	return result
}
