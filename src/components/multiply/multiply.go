package multiply

func Execute(m [][]int) (s int) {
	var response = 1
	for _, row := range m {
		for _, column := range row {
			response *= column
		}
	}
	return response
}
