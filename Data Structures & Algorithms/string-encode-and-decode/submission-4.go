type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var strCombined string

	for _, v := range strs {
		strCombined = strCombined + strconv.Itoa(len(v)) + "#" + v
	}

	return strCombined

}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	strLen := len(encoded)

	i := 0
	for i < strLen {

		j := i
		for encoded[j] != '#' {
			j++
		}

		wordLen := encoded[i:j]
		wordLenInt, _ := strconv.Atoi(wordLen)

		wordStarts := j + 1
		wordEnds := wordStarts + wordLenInt
		result = append(result, encoded[wordStarts:wordEnds])

		i = wordEnds
	}

	return result
}