func twoSum(nums []int, target int) []int {
		seen := make(map[int]int)
	result := make([]int, 0, 2)

	for i, num := range nums {
		comp := target - num
		j, ok := seen[comp]
		if ok {
			result = append(result, j)
			result = append(result, i)
		}

		seen[num] = i
	}

	return result
}
