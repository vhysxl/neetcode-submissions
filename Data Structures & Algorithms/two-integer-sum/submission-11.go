func twoSum(nums []int, target int) []int {
		seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		comp := target - v
		if _, ok := seen[comp]; ok {
			result = append(result, seen[comp])
			result = append(result, i)

			return result
		}

		seen[v] = i
	}

	return result
}
