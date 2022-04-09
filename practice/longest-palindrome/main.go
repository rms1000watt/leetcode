// https://leetcode.com/problems/longest-palindromic-substring/

package main

import (
	"fmt"

	"github.com/rms1000watt/leetcode/stringz"
)

func main() {
	fmt.Println(longestPalindrome("ijasdieeepadojasdpooooooooop"))
}

func longestPalindrome(in string) (out string) {
	longestLength := 0
	palindromes := []string{}

	if len(in) == 1 {
		out = in
		return
	}

	for i := 0; i < len(in)-1; i++ {
		for j := i + 1; j <= len(in); j++ {
			if stringz.Palindrome(in[i:j]) {
				palindromes = append(palindromes, in[i:j])
			}
		}
	}

	for _, palindrome := range palindromes {
		if len(palindrome) > longestLength {
			longestLength = len(palindrome)
			out = palindrome
		}
	}

	return
}
