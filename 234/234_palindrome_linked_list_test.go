package _234_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/234"
	"testing"
)

func Test1(t *testing.T) {
	g := NewWithT(t)

	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	g.Expect(IsPalindrome(&head)).Should(BeTrue())
}

func Test2(t *testing.T) {
	g := NewWithT(t)

	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	g.Expect(IsPalindrome(&head)).Should(BeFalse())
}
