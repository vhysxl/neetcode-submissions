func productExceptSelf(nums []int) []int {
		length := len(nums)

	prefix := make([]int, length)
	prefix[0] = 1
	for i := 1; i < length; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	suffix := make([]int, length)
	suffix[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	result := make([]int, 0, length)
	for i := range length {
		result = append(result, prefix[i]*suffix[i])
	}

	return result
}
