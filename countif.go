package piscine

func CountIf(f func(string) bool, tab []string) int {
	var list []string

	for _, str := range tab {
		if f(str) {
			list = append(list, str)
		}
	}
	return len(list)
}
