import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
		result := make([][]string, 0)
	seen := make(map[string][]string)

	for _, v := range strs {
		splitted := []byte(v)
		slices.Sort(splitted)

		seen[string(splitted)] = append(seen[string(splitted)], v)
	}

	for _, v := range seen {
		result = append(result, v)
	}

	return result
}
