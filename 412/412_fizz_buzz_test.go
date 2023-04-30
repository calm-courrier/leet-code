package _412_test

import (
	. "github.com/onsi/gomega"
	. "leet-code/412"
	"testing"
)

func Test(t *testing.T) {
	g := NewWithT(t)
	g.Expect(FizzBuzz(3)).Should(Equal([]string{"1", "2", "Fizz"}))
	g.Expect(FizzBuzz(5)).Should(Equal([]string{"1", "2", "Fizz", "4", "Buzz"}))
	g.Expect(FizzBuzz(15)).Should(Equal([]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}))
}
