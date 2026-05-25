import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
		countK := make(map[int]int)

	for _, num := range nums {

		countK[num] += 1

	}

	keyK := make([]int, 0, len(countK))
	for i := range countK {
		keyK = append(keyK, i)
	}

	slices.SortFunc(keyK, func(i, j int) int {
		return countK[j] - countK[i]
	})

	return keyK[:k]
}	
