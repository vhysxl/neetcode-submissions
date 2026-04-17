func isAnagram(s string, t string) bool {
	alphabetCount := map[string]int{}

	for _, char := range s {
		alphabetCount[string(char)] += 1
	}

	for _, char := range t {
		alphabetCount[string(char)] -= 1
	}

	for key := range alphabetCount {
		if alphabetCount[key] != 0 {
			return false
		}
	}
	
	return true
}
