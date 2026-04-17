func isPalindrome(s string) bool {
	var sliced []rune
	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			sliced = append(sliced, unicode.ToLower(v))
		}
	}

	i, j := 0, len(sliced)-1
	for i < j {
		if sliced[i] != sliced[j] {
			return false
		}

		i += 1
		j -= 1
	}

	return true
}
