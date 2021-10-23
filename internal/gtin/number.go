package gtin

import (
	"strconv"
)

type Number []int

func NewNumber(s string) Number {
	n := make(Number, len(s))

	for i, rune := range s {
		n[i], _ = strconv.Atoi(string(rune))
	}

	return n
}

func (n Number) Stringify() string {
	var s string
	for _, d := range n {
		s += strconv.Itoa(d)
	}

	return s
}
