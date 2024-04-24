package piscine

/* Function that divides a and b and stores the result
in the int pointed by div and the remainder in the int
pointed by mod */
func DivMod(a int, b int, div *int, mod *int) {

	//Declaring the variables
	if b != 0 {
		*div = a / b
		*mod = a % b
	}
}
