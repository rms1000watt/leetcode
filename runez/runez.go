package runez

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
