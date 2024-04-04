package piscine

func UltimateDivMod(a *int, b *int) {
	q := *a / *b
	r := *a % *b
	*a = q
	*b = r
}
