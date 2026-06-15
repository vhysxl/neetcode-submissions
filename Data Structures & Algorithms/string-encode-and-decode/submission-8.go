type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string

	for _, v := range strs {
		fmt.Println(v)
		result = result + strconv.Itoa(len(v)) + "*" + v
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

		wordLen := encoded[i:j]
		lenInt, _ := strconv.Atoi(wordLen)

		wordStart := j + 1
		wordEnd := wordStart + lenInt

		result = append(result, encoded[wordStart:wordEnd])

		i = wordEnd
	}

	return result
}
