func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !unicode.IsNumber(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
			continue
		}

		if !unicode.IsNumber(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j--
			continue
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true

}
