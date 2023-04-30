package _234

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	slice := make([]int, 0)
	for {
		slice = append(slice, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		if slice[i] != slice[j] {
			return false
		}
	}
	return true
}
