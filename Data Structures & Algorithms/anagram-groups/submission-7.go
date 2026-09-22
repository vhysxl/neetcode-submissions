import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
		groups := make(map[string][]string)

	for _, v := range strs {
		byteSlice := []byte(v)
		slices.Sort(byteSlice)

		groups[string(byteSlice)] = append(groups[string(byteSlice)], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}
