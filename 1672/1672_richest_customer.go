package _1672

func MaximumWealth(accounts [][]int) int {
	richest := 0
	for _, account := range accounts {
		for i, bank := range account {
			if i == 0 {
				continue
			}
			account[0] += bank
		}
		if account[0] > richest {
			richest = account[0]
		}
	}
	return richest
}
