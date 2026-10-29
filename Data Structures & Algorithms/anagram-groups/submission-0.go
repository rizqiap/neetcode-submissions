func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	var result [][]string

	for _, char := range strs {
		runes := []rune(char)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedWord := string(runes)

		groups[sortedWord] = append(groups[sortedWord], char)
	}

	for _, group := range groups {
    	result = append(result, group) 
	}

	return result
}
