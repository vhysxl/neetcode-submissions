
import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for j < k && nums[k] == nums[k-1] {
					k--
				}

				j++
				k--
			}
		}

	}

	return result
}
