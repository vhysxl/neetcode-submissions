func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		_, ok := seen[num]
		if ok {
			return true
		}
		
		seen[num] = true
	}

	return false
}
