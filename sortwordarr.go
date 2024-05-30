package piscine

// Function that sorts by ascii in ascending order a string slice.
func SortWordArr(a []string) {
	for range a {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}
