package TP

import "math"

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	targetFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetFreq[t[i]]++
	}
	left, right := 0, 0
	minLeft, minRight := 0, 0
	minLength := math.MaxInt32
	windowFreq := make(map[byte]int)
	required := len(targetFreq)
	formed := 0
	for right < len(s) {
		char := s[right]
		windowFreq[char]++
		if count, exists := targetFreq[char]; exists && windowFreq[char] == count {
			formed++
		}
		for left <= right && formed == required {
			char = s[left]
			if right-left+1 < minLength {
				minLength = right - left + 1
				minLeft = left
				minRight = right
			}
			windowFreq[char]--
			if count, exists := targetFreq[char]; exists && windowFreq[char] < count {
				formed--
			}
			left++
		}
		right++
	}
	if minLength == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minRight+1]
}
