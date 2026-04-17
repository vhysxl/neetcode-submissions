func hasDuplicate(nums []int) bool {
 	if len(nums) == 0 {
		return false
	}

	seen := map[int]bool{}

	for _, v := range nums {
		_, ok := seen[v]
		if ok {
			return true
		}

		seen[v] = true
	}

	return false
}
