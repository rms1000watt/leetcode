package stringz

// Reverse a string
func Reverse(in string) (out string) {
	for _, v := range in {
		out = string(v) + out
	}
	return
}
