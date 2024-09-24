package TP

import "sort"

func Ft_missing(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	min := nums[0]
	var missing int
	var v int
	if min != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	min = min + 1
	for v = 0; v < len(nums); v++ {
		if nums[v] == min {
			min = min + 1
			v = 0
		} else {
			missing = min
		}
	}
	return missing
}
