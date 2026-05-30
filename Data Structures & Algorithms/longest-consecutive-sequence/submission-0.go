func longestConsecutive(nums []int) int {
		isInitSeq := make(map[int]bool)
	count := 0

	for _, v := range nums {
		isInitSeq[v] = false
	}

	for _, v := range nums {
		prev := v - 1
		_, ok := isInitSeq[prev]
		if !ok {
			isInitSeq[v] = true
		}
	}

	for v, isStart := range isInitSeq {
		if isStart {
			length := 1
			for {
				_, ok := isInitSeq[v+length]
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
