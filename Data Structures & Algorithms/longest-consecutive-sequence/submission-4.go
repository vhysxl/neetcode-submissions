func longestConsecutive(nums []int) int {
		count := 0
	isInit := make(map[int]bool)

	for _, num := range nums {
		isInit[num] = false
	}

	for i := range isInit {
		_, ok := isInit[i-1]
		if !ok {
			length := 1
			for {
				_, ok := isInit[i+length]
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
