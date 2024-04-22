package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	index := ListSize(l) - 1

	curentNode := l.Tail
	previousNode := (*NodeL)(nil)
	// Finding the previous node from the current one and assigning pointer of current to the previous
	for index != 0 {
		index--
		previousNode = ListAt(l.Head, index)
		curentNode.Next = previousNode
		curentNode = previousNode
	}
	// Current is now the previous Head, so changing to point at nil,
	// changing the pointer of head to be the last element and the current to be the last.
	curentNode.Next = nil
	l.Head = l.Tail
	l.Tail = curentNode
}
