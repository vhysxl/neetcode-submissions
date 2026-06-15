func isValidSudoku(board [][]byte) bool {
		rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := range 9 {
		for column := range 9 {
			if board[row][column] == '.' {
				continue
			}

			val := board[row][column]

			box := (row/3)*3 + (column / 3)

			if rows[row][val] || columns[column][val] || boxes[box][val] {
				return false
			}

			rows[row][val] = true
			columns[column][val] = true
			boxes[box][val] = true
		}
	}

	return true
}
