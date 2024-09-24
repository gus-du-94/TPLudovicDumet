package TP

func Ft_max_substring(s string) int {
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0
	for i, char := range s {
		if lastIdx, found := charIndex[char]; found && lastIdx >= start {
			start = lastIdx + 1
		}
		charIndex[char] = i
		length := i - start + 1
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
