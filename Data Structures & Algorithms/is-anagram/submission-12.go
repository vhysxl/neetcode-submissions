func isAnagram(s string, t string) bool {
	seen := make(map[rune]int)

	for _, v := range s {
		seen[v] += 1
	}

	for _, v := range t {
		seen[v] -= 1
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}

	return true

}
