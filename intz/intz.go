package intz

// GreatestInSlice checks if 'in' is greatest in 'sl' slice
func GreatestInSlice(sl []int, in int) (out bool) {
	for _, s := range sl {
		if s > in {
			return
		}
	}

	out = true
	return
}
