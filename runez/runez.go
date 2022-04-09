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

// IncrementRuneLooped increments "inc" times from "in" and loops between "min" to "max"
func IncrementRuneLooped(min rune, max rune, in rune, inc int) (out rune) {
	out = in + rune(inc)
	overflow := out - max

	for overflow > 0 {
		out = min + overflow - 1
		overflow = out - max
	}

	return
}
