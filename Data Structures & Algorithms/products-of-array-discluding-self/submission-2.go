func productExceptSelf(nums []int) []int {
	n := len(nums)

	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	result := make([]int, n)
	result[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i+1]

	}

	for i := 0; i < n; i++ {
		result[i] *= prefix[i]
	}

	return result
}
