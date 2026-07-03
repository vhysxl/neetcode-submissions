func isAnagram(s string, t string) bool {
		alphabetCount := make(map[rune]int)

	for _, v := range s {
		alphabetCount[v] += 1
	}

	for _, v := range t {
		alphabetCount[v] -= 1
	}

	for _, v := range alphabetCount {
		if v != 0 {
			return false
		}
	}

	return true

}
