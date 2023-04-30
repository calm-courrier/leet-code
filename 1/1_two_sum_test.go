package _1_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/1"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(TwoSum([]int{2, 7, 11, 15}, 9)).Should(Equal([]int{0, 1}))
	g.Expect(TwoSum([]int{3, 2, 4}, 6)).Should(Equal([]int{1, 2}))
	g.Expect(TwoSum([]int{3, 3}, 6)).Should(Equal([]int{0, 1}))
}
