package invert

import "github.com/armanithiago/matrix-api/components/echo"

// Execute the matrix as a string in matrix format where the columns and rows are inverted
func Execute(m [][]int) (s string) {
	for r := 0; r < len(m); r++ {
		for c := 0; c < r; c++ {
			var tmp = m[r][c]
			m[r][c] = m[c][r]
			m[c][r] = tmp
		}
	}
	return echo.Execute(m)
}
