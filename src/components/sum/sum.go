package sum

// Execute returns the sum of all elements in the matrix
func Execute(m [][]int64) (s int64) {
	var response int64
	for _, row := range m {
		for _, column := range row {
			response += column
		}
	}
	return response
}
