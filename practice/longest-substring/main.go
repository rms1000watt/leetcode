// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(in string) (out int) {
	substrings := []string{}

	current := ""
	for _, c := range in {
		current += string(c)
		if ContainsDupes(current) {
			current = current[1:]
		}

		substrings = append(substrings, current)
	}

	longest := ""
	for _, substring := range substrings {
		if len(substring) > len(longest) {
			longest = substring
		}
	}

	out = len(longest)
	return
}

func ContainsDupes(in string) (out bool) {
	for _, c := range in {
		if strings.Count(in, string(c)) > 1 {
			out = true
			return
		}
	}

	return
}
