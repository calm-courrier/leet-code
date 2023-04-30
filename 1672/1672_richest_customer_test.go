package _1672_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/1672"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}})).Should(Equal(6))
	g.Expect(MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})).Should(Equal(10))
	g.Expect(MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})).Should(Equal(17))
}
