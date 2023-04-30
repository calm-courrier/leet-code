package _1342

func NumberOfSteps(num int) int {
	steps := 0
	for {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}
