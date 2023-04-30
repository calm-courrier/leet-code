package _1929_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/1929"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(GetConcatenation([]int{1, 2, 1})).Should(Equal([]int{1, 2, 1, 1, 2, 1}))
	g.Expect(GetConcatenation([]int{1, 3, 2, 1})).Should(Equal([]int{1, 3, 2, 1, 1, 3, 2, 1}))
}
