type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string
	for _, str := range strs {
		result = result + strconv.Itoa(len(str)) + "*" + str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	length := len(encoded)
	var result []string

	i := 0
	for i < length {
		j := i
		for encoded[j] != '*' {
			j++
		}

		wordLength := encoded[i:j]
		lengthInt, _ := strconv.Atoi(string(wordLength))

		wordStart := j + 1
		wordEnd := wordStart + lengthInt
		result = append(result, encoded[wordStart:wordEnd])

		i = wordEnd

	}

	return result

}
