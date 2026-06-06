func longestConsecutive(nums []int) int {
	count := 0

	numsMap := make(map[int]bool)
	for _, v := range nums {
		numsMap[v] = false
	}

	for i := range numsMap {
		if _, ok := numsMap[i-1]; !ok {
			length := 1

			for {
				_, ok := numsMap[i+length]
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
