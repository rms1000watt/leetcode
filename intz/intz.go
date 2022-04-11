package intz

import (
	"errors"
)

// GreatestInSlice checks if 'in' is greatest in 'sl' slice
func GreatestInSlice(sl []int, in int) (out bool) {
	for _, s := range sl {
		if s > in {
			return
		}
	}

	out = true
	return
}

// ReversePositive reverses a positive integer like 857 to 758
func ReversePositive(in int) (out int, err error) {
	if in < 0 {
		err = errors.New("this function only accepts positive integers")
		return
	}

	for in > 0 {
		mod := in % 10
		out = out*10 + mod
		in /= 10
	}

	return
}

// ReverseNegative reverses a negative integer like -857 to -758
func ReverseNegative(in int) (out int, err error) {
	if in > 0 {
		err = errors.New("this function only accepts positive integers")
		return
	}

	out, err = ReversePositive(-1 * in)
	out = -1 * out

	return
}
