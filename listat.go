package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}

	currentNode := l
	index := 0

	for currentNode != nil {
		if pos == index {
			return currentNode
		}
		currentNode = currentNode.Next
		index++
	}
	return nil
}
