func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))
	x := make([][]int, 0, len(nums))

	for i := range nums {
		y := append([]int(nil), nums[:i]...)
		y = append(y, nums[i+1:]...)
		x = append(x, y)
	}

	for _, v := range x {
		res := 1

		for _, num := range v {
			res = res * num
		}

		result = append(result, res)
	}

	return result
}
