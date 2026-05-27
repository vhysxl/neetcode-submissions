func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		compl := target - v
		if x, ok := seen[compl]; ok {
			result = append(result, x)
			result = append(result, i)

			return result
		}

		seen[v] = i
	}

	return result
}
