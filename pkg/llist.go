package main

type LlNode[T any] struct {
	Next *LlNode[T]
	Prev *LlNode[T]
	Val T
}

type Llist[T any] struct {
	Start *LlNode[T]
	End *LlNode[T]
}

func (l *Llist[T]) Append(vals ...T) []*LlNode[T] {
	var nodes []*LlNode[T]
	for _, v := range vals {
		node := new(LlNode[T])
		nodes = append(nodes, node)
		node.Prev = l.End
		node.Val = v
		if l.Start == nil {
			l.Start = node
		} else if l.End != nil {
			l.End.Next = node
		}
		l.End = node
	}
	return nodes
}

func (l *Llist[T]) Del(nodes ...*LlNode[T]) {
	for _, n := range nodes {
		if n.Prev == nil {
			l.Start = n.Next
		} else {
			n.Prev.Next = n.Next
		}
		if n.Next == nil {
			l.End = n.Prev
		} else {
			n.Next.Prev = n.Prev
		}
	}
}

