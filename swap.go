package piscine

/* Function that takes two pointers to an int
and swaps their contents */
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
