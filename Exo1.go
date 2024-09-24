package TP

func Ft_coin(coins []int, amount int) int {
	var max = coins[0]
	var t int
	var v int
	for amount != 0 {
		if amount == 1 {
			for v = 0; v < len(coins); v++ {
				if coins[v] == 1 {
					amount = amount - 1
				} else {
					return -1
				}
			}
		} else {
			for i := 1; i < len(coins); i++ {
				if max < coins[i] && coins[i] < amount {
					max = coins[i]
				}
			}
		}
		amount = amount - max
		max = coins[0]
		t = t + 1
	}
	return t
}
