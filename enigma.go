package piscine

// Function that receives pointers as arguments and move its values
// around to hide them. This function will put:
// a into c, c into d,
// d into b, b into a.
func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Change all the values at the same time
	*******c, ****d, ***a, *b = ***a, *******c, *b, ****d
}
