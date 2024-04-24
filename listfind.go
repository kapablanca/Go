package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	currentNode := l.Head

	for currentNode != nil {
		if comp(currentNode.Data, ref) {
			return &currentNode.Data
		}
		currentNode = currentNode.Next
	}
	return nil
}
