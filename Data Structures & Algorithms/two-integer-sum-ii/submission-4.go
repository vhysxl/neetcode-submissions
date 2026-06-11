func twoSum(numbers []int, target int) []int {
    	result := make([]int, 2)

	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			result[0] = i + 1
			result[1] = j + 1
			break
		}
	}

	return result
}
