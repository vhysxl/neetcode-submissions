func isPalindrome(s string) bool {
	sliced := make([]rune, 0, len(s))

	for _, char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			sliced = append(sliced, unicode.ToLower(char))
		}
	}

	i, j := 0, len(sliced)-1
	for i < j {
		if sliced[i] != sliced[j] {
			return false
		}

		i++
		j--
	}


	return true
}
