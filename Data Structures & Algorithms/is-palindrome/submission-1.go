func isPalindrome(s string) bool {
	sliced := make([]rune, 0, len(s))
	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			sliced = append(sliced, unicode.ToLower(v))
		}
	}

	clone := make([]rune, len(sliced))
	copy(clone, sliced)

	for i, j := 0, len(sliced)-1; i < j; i, j = i+1, j-1 {
		sliced[i], sliced[j] = sliced[j], sliced[i]
	}

	if string(clone) == string(sliced) {
		return true
	}

	return false
}
