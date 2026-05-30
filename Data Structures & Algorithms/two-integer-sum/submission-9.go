func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		comp := target - v
		if j, ok := seen[comp]; ok {
			result = append(result, j)
			result = append(result, i)
		}

		seen[v] = i
	}

	return result
}
