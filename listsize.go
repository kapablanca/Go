package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	size := 1
	if l.Head == nil {
		size = 0
		return size
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			size++
			currentNode = currentNode.Next
		}
	}
	return size
}
