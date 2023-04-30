package _1

func TwoSum(nums []int, target int) []int {
	for i, numA := range nums {
		for j, numB := range nums {
			if i == j {
				continue
			}
			if numA+numB == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
