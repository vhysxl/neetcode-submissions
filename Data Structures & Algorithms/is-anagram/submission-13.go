func isAnagram(s string, t string) bool {
	letter := make(map[rune]int)

	for _, v := range s {
		letter[v] += 1
	}

	for _, v := range t {
		letter[v] -= 1
	}

	for _, v := range letter {
		if v != 0 {
			return false
		}
	}

	return true
}
