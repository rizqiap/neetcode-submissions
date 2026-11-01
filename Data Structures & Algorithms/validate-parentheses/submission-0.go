func isValid(s string) bool {
    stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, b := range s {
		if open, exists := closeToOpen[b]; exists {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if  top != open {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
