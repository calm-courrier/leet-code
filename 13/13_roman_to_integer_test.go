package _13_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/13"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(RomanToInt("V")).Should(Equal(5))
	g.Expect(RomanToInt("III")).Should(Equal(3))
	g.Expect(RomanToInt("LVIII")).Should(Equal(58))
	g.Expect(RomanToInt("D")).Should(Equal(500))
	g.Expect(RomanToInt("MCMXCIV")).Should(Equal(1994))
	g.Expect(func() { RomanToInt("MCMA") }).Should(Panic())
}
