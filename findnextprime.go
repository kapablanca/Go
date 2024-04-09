package piscine

func FindNextPrime(nb int) int {
	const MaxInt = 2147483647

	for i := nb; i < MaxInt; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
