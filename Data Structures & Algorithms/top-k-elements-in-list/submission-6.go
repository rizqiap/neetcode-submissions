func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
    }

	keys := make([]int, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool{
		return freq[keys[i]] > freq[keys[j]]
	})

	res := make([]int, k)
	for i := 0; i < k ; i++ {
		res[i] = keys[i]
	}

    return res
}
