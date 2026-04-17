func isAnagram(s string, t string) bool {
	seen := map[string]int{}

	for _, v := range s {
		seen[string(v)] += 1
	}

	for _, v := range t {
		seen[(string(v))] -= 1
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}

	return true
}
