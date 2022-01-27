package flatten

import (
	"fmt"
	"strings"
)

// Execute returns the matrix as a 1 line string, with values separated by commas.
func Execute(m [][]int) (s string) {
	var result = ""
	for _, row := range m {
		if result != "" {
			result += ","
		}
		strRow := strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ","), "[]")
		result = fmt.Sprintf("%s%s", result, strRow)
	}
	return result
}
