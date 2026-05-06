func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, x := range nums {
		comp := target - x

		res, ok := seen[comp]
		if ok {
			result = append(result, res)
			result = append(result, i)
		}

		seen[x] = i

	}

	return result
}
