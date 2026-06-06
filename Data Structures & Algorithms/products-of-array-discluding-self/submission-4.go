func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	for i := range len(nums) {
		numsToProduct := make([]int, 0, length-1)

		for j, v := range nums {
			if j == i {
				continue
			}

			numsToProduct = append(numsToProduct, v)
		}

		product := 1
		for _, v := range numsToProduct {
			product *= v

		}

		result[i] = product

	}

	return result
}
