package piscine

func StrRev(s string) string {
	aString := []byte(s)

	length := len(aString)

	for i := 0; i < length/2; i++ {
		aString[i], aString[length-i-1] = aString[length-i-1], aString[i]
	}

	return string(aString)
}
