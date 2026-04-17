func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	cleaned := make([]rune, 0, len(s))
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			cleaned = append(cleaned, v)
		}
	}

	cleanedJoined := string(cleaned)

	slicedReversed := strings.Split(cleanedJoined, "")

	for i, j := 0, len(slicedReversed)-1; i < j; i, j = i+1, j-1 {
		slicedReversed[i], slicedReversed[j] = slicedReversed[j], slicedReversed[i]
	}

	finalReversed := strings.Join(slicedReversed, "")

	if finalReversed == cleanedJoined {
		return true
	}

	return false
}
