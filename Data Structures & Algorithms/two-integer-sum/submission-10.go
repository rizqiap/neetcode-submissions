func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, num := range nums {
		compliment := target - num 

		index, exist := seen[compliment]

		if exist {
			return []int{index, i}
		} else {
			seen[num] = i
		}
	}

	return nil
}
