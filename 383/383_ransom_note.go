package _383

func CanConstruct(ransomNote string, magazine string) bool {
	magazineMap := map[rune]int{}
	for _, c := range magazine {
		magazineMap[c] += 1
	}
	for _, c := range ransomNote {
		magazineMap[c] -= 1
		if magazineMap[c] < 0 {
			return false
		}
	}
	return true
}
