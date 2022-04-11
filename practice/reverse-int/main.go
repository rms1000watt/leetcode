// https://leetcode.com/problems/reverse-integer/

package main

import (
	"fmt"
	"math"

	"github.com/rms1000watt/leetcode/intz"
)

func main() {
	fmt.Println(reverse(4681))
	fmt.Println(reverse(-4681))
	fmt.Println(reverse(1534236469))
}

func reverse(in int) (out int) {
	if in > 0 {
		out, _ = intz.ReversePositive(in)
		out = outsideBounds(out)
		return
	}

	if in < 0 {
		out, _ = intz.ReverseNegative(in)
		out = outsideBounds(out)
		return
	}

	return
}

func outsideBounds(in int) (out int) {
	if in > int(math.Pow(2, 31))-1 {
		return
	}

	if in < -1*int(math.Pow(2, 31)) {
		return
	}

	out = in
	return
}
