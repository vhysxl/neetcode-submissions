func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
		}
		fmt.Println(i)

		for i < j && !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j--
		}
		fmt.Println(j)

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
