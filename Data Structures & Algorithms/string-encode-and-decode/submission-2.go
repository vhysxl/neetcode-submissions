type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var result string
	separator := "*"

	for _, str := range strs {
		strLen := len(str)
		result = result + strconv.Itoa(strLen) + separator + str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	lenStr := len(encoded)
	i := 0

	for i < lenStr {

		j := i
		for encoded[j] != '*' {
			j++
		}

		wordLen := encoded[i:j]
		fmt.Println(wordLen)
		lenInt, _ := strconv.Atoi(wordLen)

		startWord := j + 1
		endWord := startWord + lenInt

		result = append(result, encoded[startWord:endWord])

		i = endWord
	}

	return result
}