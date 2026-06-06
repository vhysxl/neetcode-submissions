func isPalindrome(s string) bool {
	
	var validChar []string
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			validChar = append(validChar, string(unicode.ToLower(v)))
		}
	}

	for i, j := 0, len(validChar)-1; i < j; {
		if validChar[i] != validChar[j] {
			return false
		}

		i++
		j--
	}

	return true
}

