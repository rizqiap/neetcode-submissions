func findMin(nums []int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		m := l + (r - l)/2

		fmt.Println("mid: ", m)

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[r]
}

	

