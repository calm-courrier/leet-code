package _876_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/876"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	g.Expect(MiddleNode(&node).Val).Should(Equal(3))
	node = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	g.Expect(MiddleNode(&node).Val).Should(Equal(4))
}
