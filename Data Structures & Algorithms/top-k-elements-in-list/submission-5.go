import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, v := range nums {
		seen[v] += 1
	}

	result := make([]int, 0, len(seen))
	for i := range seen {
		result = append(result, i)
	}

	slices.SortFunc(result, func(i, j int) int {
		return seen[j] - seen[i]
	})

	return result[:k]
}	
