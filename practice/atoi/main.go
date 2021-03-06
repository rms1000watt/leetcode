// https://leetcode.com/problems/string-to-integer-atoi/

package main

import (
	"fmt"
	"math"
)

func main() {
	inputs := []string{
		" 9129as-di ",
		"-42",
		"0032",
		"words and 987",
		"9223372036854775808",
	}

	for _, input := range inputs {
		fmt.Println(myAtoi(input))
	}
}

const (
	stateTrimLeft     = "TRIMLEFT"
	stateReadIntegers = "READINTEGERS"
	signPositive      = "POSITIVE"
	signNegative      = "NEGATIVE"
)

func myAtoi(in string) (out int) {
	integers := Slice('0', '9')
	state := stateTrimLeft
	sign := signPositive

	for _, c := range in {
		// fmt.Println("state:", state)
		// fmt.Println("sign:", sign)

		if state == stateTrimLeft {
			if c == ' ' {
				continue
			}

			if c != '-' && c != '+' && !Contains(integers, c) {
				return
			}

			if c == '+' {
				sign = signPositive
				state = stateReadIntegers
				continue
			}

			if c == '-' {
				sign = signNegative
				state = stateReadIntegers
				continue
			}

			if Contains(integers, c) {
				state = stateReadIntegers
			}
		}

		if state == stateReadIntegers {
			if !Contains(integers, c) {
				break
			}

			out = out*10 + int(c-'0')

			if out > int(math.Pow(2, 31)) {
				out = int(math.Pow(2, 31))
			}
		}
	}

	if sign == signNegative {
		out = -1 * out
	}

	if out > int(math.Pow(2, 31))-1 {
		out = int(math.Pow(2, 31)) - 1
		return
	}

	if out < -1*int(math.Pow(2, 31)) {
		out = -1 * int(math.Pow(2, 31))
		return
	}

	return
}

// Slice creates a rune slice from start to end rune
func Slice(start rune, end rune) (out []rune) {
	for c := start; c <= end; c++ {
		out = append(out, c)
	}
	return
}

// Contains check if the runeSlice contains a rune
func Contains(runeSlice []rune, check rune) (out bool) {
	for _, c := range runeSlice {
		if c == check {
			out = true
			return
		}
	}
	return
}
