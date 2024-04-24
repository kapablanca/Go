package piscine

/* Function that takes as argument a pointer
 to a pointer to a pointer to an integer
and gives it the value of 1 */
func UltimatePointOne(n ***int) {
	***n = 1
}
