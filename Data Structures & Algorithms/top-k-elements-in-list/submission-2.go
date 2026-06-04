import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
		count := make(map[int]int)

	for _, v := range nums {
		count[v] += 1
	}

	result := make([]int, 0, k)
	for i := range count {
		result = append(result, i)
	}

	slices.SortFunc(result, func(a, b int) int {
		return count[b] - count[a]
	})

	return result[:k]
}	
