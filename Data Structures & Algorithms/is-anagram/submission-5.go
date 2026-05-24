func isAnagram(s string, t string) bool {
	alphabetCount := make(map[string]int)

	for _, v := range s {
		alphabetCount[string(v)] += 1
	}

	for _, v := range t {
		alphabetCount[string(v)] -= 1
	}

	for _, v := range alphabetCount {
		if v != 0 {
			return false
		}
	}

	return true
}
