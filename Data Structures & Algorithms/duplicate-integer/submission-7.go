func hasDuplicate(nums []int) bool {

	seen := make(map[int]bool)

	for _, val := range nums {
		_, ok := seen[val]
		if ok {
			return true
		}

		seen[val] = true
	}

	return false
}
