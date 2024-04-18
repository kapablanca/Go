package piscine

func Abort(a, b, c, d, e int) int {
	list := [5]int{a, b, c, d, e}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list[2]
}
