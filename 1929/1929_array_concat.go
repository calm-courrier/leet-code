package _1929

func GetConcatenation(nums []int) []int {
	l := len(nums)
	res := make([]int, l*2)
	for i := 0; i < l*2; i++ {
		res[i] = nums[i%l]
	}
	return res
}
