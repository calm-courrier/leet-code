package _1337

func KWeakestRows(mat [][]int, k int) []int {
	freq := make([][]int, len(mat[0])+1)
	for i := range freq {
		freq[i] = make([]int, 0)
	}
	for rowI, row := range mat {
		soldierCount := 0
		for _, val := range row {
			if val != 1 {
				break
			}
			soldierCount++
		}
		freq[soldierCount] = append(freq[soldierCount], rowI)
	}
	res := make([]int, k)
	resI := 0
outer:
	for i := 0; i < len(freq); i++ {
		for _, val := range freq[i] {
			res[resI] = val
			resI++
			if resI == k {
				break outer
			}
		}
	}
	return res
}
