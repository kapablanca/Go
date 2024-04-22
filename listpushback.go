package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	nextNode := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = nextNode
		l.Tail = nextNode

	} else {
		l.Tail.Next = nextNode
	}
	l.Tail = nextNode
}
