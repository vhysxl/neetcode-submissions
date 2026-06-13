import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
		groups := make(map[string][]string)

	for _, str := range strs {
		bit := []byte(str)
		slices.Sort(bit)
		word := string(bit)

		_, ok := groups[word]
		if ok {
			groups[word] = append(groups[word], str)
			continue
		}
		groups[word] = append(groups[word], str)
	}

	result := make([][]string, 0, len(groups))

	for _, val := range groups {
		result = append(result, val)
	}

	return result
}
