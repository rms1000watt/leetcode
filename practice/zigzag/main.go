// https://leetcode.com/problems/zigzag-conversion/submissions/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("HELLOWORLD", 2))
}

func convert(in string, numRows int) (out string) {
	if len(in) == 1 || numRows == 1 {
		out = in
		return
	}

	matrix := make([][]string, numRows)

	row := 0
	zig := true
	for _, s := range in {
		matrix[row] = append(matrix[row], string(s))

		if zig {
			row++
		} else {
			row--
		}

		if row >= numRows {
			row -= 2
			zig = false
		}

		if row == -1 {
			row += 2
			zig = true
		}
	}

	for _, row := range matrix {
		for _, s := range row {
			out += string(s)
		}
	}

	return
}
