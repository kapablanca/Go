package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	nextNode := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Tail == nil {
		l.Tail = nextNode
		l.Head = nextNode
	} else {
		nextNode.Next = l.Head
	}
	l.Head = nextNode
}
