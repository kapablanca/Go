package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListClear(l *List) {
	var tempNode *NodeL

	for l.Head != nil {
		tempNode = l.Head.Next
		l.Head = nil
		l.Head = tempNode
	}
}
