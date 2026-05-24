import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, v := range nums {
		count[v] += 1
	}

	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(i, j int) int { return count[j] - count[i] })

	return keys[:k]
}	
