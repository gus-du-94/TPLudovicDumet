package TP

func Ft_profit(prices []int) int {
	min := prices[0]
	var benef int
	var i int
	index := 0
	for i = 1; i < (len(prices) - 1); i++ {
		if min < prices[i] {
			continue
		} else {
			min = prices[i]
			index = i
		}
	}
	max := prices[index]
	for v := index + 1; v < (len(prices)); v++ {
		if max < prices[v] {
			max = prices[v]
		} else {
			continue
		}
	}
	benef = max - min
	if benef < 0 {
		benef = 0
	}
	return benef
}
