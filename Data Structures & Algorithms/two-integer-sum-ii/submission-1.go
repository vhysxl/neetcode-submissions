func twoSum(numbers []int, target int) []int {
    	result := make([]int, 0, 2)

	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			result = append(result, i+1)
			result = append(result, j+1)

			break
		}
	}

	return result
}
