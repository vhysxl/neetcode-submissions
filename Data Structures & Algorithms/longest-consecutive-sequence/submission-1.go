func longestConsecutive(nums []int) int {
	count := 0
	isInitialNums := make(map[int]bool)

	for _, v := range nums {
		isInitialNums[v] = false
	}

	for i := range isInitialNums {
		if _, ok := isInitialNums[i-1]; !ok {
			length := 1
			for {
				_, ok := isInitialNums[i+length]
				if !ok {
					break
				}
				length++
			}

			if length > count {
				count = length
			}
		}
	}

	return count
}
