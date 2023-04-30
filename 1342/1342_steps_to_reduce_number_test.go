package _1342_test

import (
	. "github.com/onsi/gomega"
	_1342 "leet-code/1342"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(_1342.NumberOfSteps(14)).Should(Equal(6))
	g.Expect(_1342.NumberOfSteps(8)).Should(Equal(4))
	g.Expect(_1342.NumberOfSteps(123)).Should(Equal(12))
}
