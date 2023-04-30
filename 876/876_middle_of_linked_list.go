package _876

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	slice := make([]*ListNode, 0)
	for {
		slice = append(slice, head)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	length := len(slice)
	return slice[(length / 2)]
}
