func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		comp := target - v
		if val, ok := seen[comp]; ok {
			result = append(result, val)
			result = append(result, i)
		}

		seen[v] = i
	}

	return result
}
