package runez

// Slice creates a rune slice from start to end rune
func Slice(start rune, end rune) (out []rune) {
	for c := start; c <= end; c++ {
		out = append(out, c)
	}
	return
}
