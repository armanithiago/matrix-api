package sum

// Execute returns the sum of all elements in the matrix
func Execute(m [][]int) (s int) {
	var response int
	for _, row := range m {
		for _, column := range row {
			response += column
		}
	}
	return response
}
