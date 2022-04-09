package stringz

// Reverse a string
func Reverse(in string) (out string) {
	for _, v := range in {
		out = string(v) + out
	}
	return
}

// Palindrome checks if string is a palindrome
func Palindrome(in string) (out bool) {
	if len(in) == 1 {
		out = true
		return
	}

	for i := 0; i < (len(in)/2)+1; i++ {
		j := len(in) - 1 - i
		if in[i] != in[j] {
			return
		}
	}

	out = true
	return
}
