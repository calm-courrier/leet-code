package _383_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/383"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(CanConstruct("a", "b")).Should(BeFalse())
	g.Expect(CanConstruct("aa", "ab")).Should(BeFalse())
	g.Expect(CanConstruct("aa", "aab")).Should(BeTrue())
}
