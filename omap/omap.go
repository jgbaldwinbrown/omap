package main

import (
	"fmt"
)

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

type OmapEntry[Key comparable, Value any] struct {
	Node *LlNode[Key]
	Val Value
}

type Omap[Key comparable, Value any] struct {
	List Llist[Key]
	Map map[Key]OmapEntry[Key, Value]
}

func NewOmap[Key comparable, Value any]() *Omap[Key, Value] {
	out := new(Omap[Key, Value])
	out.Map = make(map[Key]OmapEntry[Key, Value])
	return out
}

func (o *Omap[Key, Value]) Set(k Key, v Value) {
	var node *LlNode[Key]
	entry, present := o.Map[k]
	if !present {
		nodes := o.List.Append(k)
		node = nodes[0]
	} else {
		node = entry.Node
	}
	o.Map[k] = OmapEntry[Key, Value]{node, v}
}

func (o *Omap[Key, Value]) Get(k Key) (Value, bool) {
	var zero Value
	entry, ok := o.Map[k]
	if !ok {
		return zero, false
	}
	return entry.Val, true
}

func (o *Omap[Key, Value]) Del(k Key) bool {
	entry, ok := o.Map[k]
	if ok {
		node := entry.Node
		o.List.Del(node)
		delete(o.Map, k)
	}
	return ok
}

type Orange[Key comparable, Value any] struct {
	Omap *Omap[Key, Value]
	Started bool
	Current *LlNode[Key]
}

func (o *Omap[Key, Value]) Range() Orange[Key, Value] {
	return Orange[Key, Value]{o, false, nil}
}

func (o *Orange[Key, Value]) Next() bool {
	if !o.Started {
		o.Started = true
		o.Current = o.Omap.List.Start
	} else {
		o.Current = o.Current.Next
	}
	return o.Current != nil
}

func (o *Orange[TKey, TValue]) Key() (TKey, bool) {
	var zero TKey
	if o.Current == nil {
		return zero, false
	}
	return o.Current.Val, true
}

func (o *Orange[TKey, TValue]) Val() (TValue, bool) {
	var zero TValue
	if o.Current == nil {
		return zero, false
	}
	return o.Omap.Map[o.Current.Val].Val, true
}

func main() {
	m := NewOmap[string,float64]()
	m.Set("apple", 1.0)
	m.Set("banana", 2.0)
	fmt.Println(m)
	m.Set("apple", 3.0)
	m.Del("banana")
	v, ok := m.Get("apple")
	fmt.Println(v, ok)
	fmt.Println(m)
	v, ok = m.Get("banana")
	fmt.Println(v, ok)

	m.Set("earlobe", 99.9)
	m.Set("hat", 7)
	fmt.Println(m)
	r := m.Range()
	for r.Next() {
		key, kok := r.Key()
		val, vok := r.Val()
		fmt.Println(key, val, kok, vok)
	}
}
