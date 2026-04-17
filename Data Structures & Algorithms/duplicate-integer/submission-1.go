func hasDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	seen := map[int]bool{}
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}

		seen[v] = true
	}

	return false
}
