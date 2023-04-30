package _1337_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/1337"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	g.Expect(KWeakestRows(mat, 3)).Should(Equal([]int{2, 0, 3}))

	mat = [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	g.Expect(KWeakestRows(mat, 2)).Should(Equal([]int{0, 2}))
}
