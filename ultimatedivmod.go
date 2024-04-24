package piscine

/* Function that divides the dereferenced value of a
by the dereferenced value of b, stores the result of
the division in the int which a points to and stores
the remainder of the division in the int which b points to */
func UltimateDivMod(a *int, b *int) {
	if *b != 0 {
		// Stores temporary the dereferenced value of a
		temp := *a
		*a = *a / *b
		*b = temp % *b
	}
}
