func isAnagram(s string, t string) bool {
	charCount := make(map[string]int)

	for _, v := range s {
		charCount[string(v)] += 1
	}

	for _, v := range t {
		charCount[string(v)] -= 1
	}

	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}

	return true
}
