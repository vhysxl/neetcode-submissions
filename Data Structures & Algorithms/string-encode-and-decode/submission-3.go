type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var result string

	for _, v := range strs {
		result = result + strconv.Itoa(len(v)) + "#" + v
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	lenStr := len(encoded)
	var result []string
	i := 0

	for i < lenStr {
		j := i

		for encoded[j] != '#' {
			j++
		}

		wordLen := encoded[i:j]
		wordLenInt, _ := strconv.Atoi(wordLen)

		wordStartIndex := j + 1
		wordEndIndex := wordStartIndex + wordLenInt
		word := encoded[wordStartIndex:wordEndIndex]

		result = append(result, word)

		i = wordEndIndex

	}

	return result
}