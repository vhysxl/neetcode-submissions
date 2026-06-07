func twoSum(numbers []int, target int) []int {
 	var x, y int

	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			x = i + 1
			y = j + 1
			break
		}

	}

	return []int{x, y}
}   
