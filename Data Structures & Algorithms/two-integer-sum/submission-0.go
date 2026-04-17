func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		compl := target - v
		val, ok := seen[compl]
		if ok {
			result = append(result, val)
			result = append(result, i)
		}

		seen[v] = i
	}

	return result
}
