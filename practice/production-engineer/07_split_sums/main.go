package main

import "fmt"

type Match struct {
	Left  []int
	Right []int
	Sum   int
}

func main() {
	in1 := []int{2, 3, 4, 5}
	in2 := []int{2, 3, 4, 9}

	splitSums(in1)
	splitSums(in2)
}

func splitSums(in []int) (out bool) {
	sum := 0
	for _, v := range in {
		sum += v
	}

	if sum%2 == 1 {
		return
	}

	halfSum := sum / 2

	matches := []Match{}

	for i := 0; i < len(in); i++ {
		leftSum := 0
		left := []int{}
		for j := 0; j < i; j++ {
			leftSum += in[j]
			left = append(left, in[j])
		}

		if leftSum > halfSum {
			continue
		}

		rightSum := 0
		right := []int{}
		for j := i; j < len(in); j++ {
			rightSum += in[j]
			right = append(right, in[j])
		}

		if rightSum > halfSum {
			continue
		}

		if rightSum == leftSum {
			matches = append(matches, Match{
				Left:  left,
				Right: right,
				Sum:   rightSum,
			})
		}

		// fmt.Println(left)
		// fmt.Println(right)
	}

	fmt.Println(matches)
	return
}
