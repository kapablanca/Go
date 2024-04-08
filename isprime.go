package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	root := Sqrt(nb)
	for i := 2; i <= root; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
