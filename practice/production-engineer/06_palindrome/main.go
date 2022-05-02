package main

import "fmt"

func main() {
	in1 := "ijasidja"
	in2 := "racecar"
	in3 := ""
	in4 := "aa"

	fmt.Println(isPalindrome(in1))
	fmt.Println(isPalindrome(in2))
	fmt.Println(isPalindrome(in3))
	fmt.Println(isPalindrome(in4))
}

func isPalindrome(in string) (out bool) {
	if len(in) == 0 {
		return
	}

	if len(in) == 1 {
		out = true
		return
	}

	for i := 0; i < len(in)/2+1; i++ {
		j := len(in) - 1 - i
		if in[i] != in[j] {
			return
		}
	}

	out = true
	return
}
