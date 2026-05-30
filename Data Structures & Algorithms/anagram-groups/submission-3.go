import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	seen := make(map[string][]string)

	for _, v := range strs {
		b := []byte(v)
		slices.Sort(b)

		seen[string(b)] = append(seen[string(b)], v)
	}

	result := make([][]string, 0, len(seen))
	for _, v := range seen {
		result = append(result, v)
	}

	return result
}
