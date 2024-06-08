package piscine

// Function that returns the number of active bits (bits with the value of 1)
//  in the binary representation of an integer number.
func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		// Check the least significand bit with bitwise AND
		if n&1 == 1 {
			count += 1
		}
		// Right shift the bits of n
		n >>= 1
	}
	return count
}
