func isValidSudoku(board [][]byte) bool {

	column := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range 9 {
		column[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := range 9 {
		for c := range 9 {
			val := board[r][c]

			if val == '.' {
				continue
			}

			box := (r/3)*3 + (c / 3)

			if rows[r][val] || column[c][val] || boxes[box][val] {
				return false
			}

			rows[r][val] = true
			column[c][val] = true
			boxes[box][val] = true
		}
	}

	return true
}
