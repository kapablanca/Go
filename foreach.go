package piscine

// Function that for an int slice, applies a function
// on each element of that slice

func ForEach(f func(int), a []int) {
	for _, elem := range a {
		f(elem)
	}
}
