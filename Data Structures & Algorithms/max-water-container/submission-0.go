func maxArea(heights []int) int {
		max := 0

	for i, j := 0, len(heights)-1; i < j; {
		length := j - i
		minHeight := min(heights[i], heights[j])

		total := length * minHeight
		if total > max {
			max = total
		}

		if minHeight == heights[i] {
			i++
			continue
		} else {
			j--
			continue
		}
	}

	return max
}
