package _13

import (
	"errors"
	"fmt"
)

func RomanToInt(s string) int {
	var numArr []int
	for _, c := range s {
		toInt, err := convertToInt(c)
		if err != nil {
			panic(err)
		}
		numArr = append(numArr, toInt)
	}
	var sum = 0
	var prevVal = 0
	for _, num := range numArr {
		sum += num
		if num > prevVal {
			sum -= prevVal * 2
		}
		prevVal = num
	}

	return sum
}

func convertToInt(c rune) (int, error) {
	var val int
	var err error
	switch c {
	case 'I':
		val = 1
	case 'V':
		val = 5
	case 'X':
		val = 10
	case 'L':
		val = 50
	case 'C':
		val = 100
	case 'D':
		val = 500
	case 'M':
		val = 1000
	default:
		err = errors.New(fmt.Sprintf("%c is not valid Roman character", c))
	}
	return val, err
}
