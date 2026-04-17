func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	result := make([]int, 0, 2)

	for i, v := range nums {
		comp := target - v
		val, exist := seen[comp]
		if exist {
			result = append(result, val)
			result = append(result, i)
		}

		seen[v] = i
	}

	return result
}
