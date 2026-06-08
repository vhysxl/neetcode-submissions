
import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	seen := make(map[[3]int]bool)

	slices.Sort(nums)

	for i := range nums {
		for j, k := 0, len(nums)-1; j < k; {
			if i == j {
				j++
				continue
			} else if i == k {
				k--
				continue
			}

			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				val := [3]int{nums[i], nums[j], nums[k]}
				slices.Sort(val[:])
				if !seen[val] {
					result = append(result, []int{val[0], val[1], val[2]})
					seen[val] = true
				}
				j++
			}

		}
	}

	return result
}
