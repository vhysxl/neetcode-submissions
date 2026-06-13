import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	topK := make(map[int]int)

	for _, num := range nums {
		topK[num] += 1
	}

	result := make([]int, 0, len(topK))
	for x := range topK {
		result = append(result, x)
	}

	slices.SortFunc(result, func(i, j int) int {
		return topK[j] - topK[i]
	})

	return result[:k]
}	
