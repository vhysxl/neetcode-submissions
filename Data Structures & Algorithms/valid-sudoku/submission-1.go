func isValidSudoku(board [][]byte) bool {

	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := range 9 {
		for c := range 9 {

			val := board[r][c]

			if val == '.' {
				continue
			}

			box := (r/3)*3 + (c / 3)

			if rows[r][val] || columns[c][val] || boxes[box][val] {
				return false
			}

			rows[r][val] = true
			columns[c][val] = true
			boxes[box][val] = true
		}
	}

	return true

}
