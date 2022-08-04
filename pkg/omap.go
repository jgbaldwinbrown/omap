package omap

type MapEntry[Key comparable, Value any] struct {
	Node *LlNode[Key]
	Val Value
}

type Map[Key comparable, Value any] struct {
	List Llist[Key]
	Map map[Key]MapEntry[Key, Value]
}

func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	out := new(Map[Key, Value])
	out.Map = make(map[Key]MapEntry[Key, Value])
	return out
}

func (o *Map[Key, Value]) Set(k Key, v Value) {
	var node *LlNode[Key]
	entry, present := o.Map[k]
	if !present {
		nodes := o.List.Append(k)
		node = nodes[0]
	} else {
		node = entry.Node
	}
	o.Map[k] = MapEntry[Key, Value]{node, v}
}

func (o *Map[Key, Value]) Get(k Key) (Value, bool) {
	var zero Value
	entry, ok := o.Map[k]
	if !ok {
		return zero, false
	}
	return entry.Val, true
}

func (o *Map[Key, Value]) Del(k Key) bool {
	entry, ok := o.Map[k]
	if ok {
		node := entry.Node
		o.List.Del(node)
		delete(o.Map, k)
	}
	return ok
}

type Range[Key comparable, Value any] struct {
	Map *Map[Key, Value]
	Started bool
	Current *LlNode[Key]
}

func (o *Map[Key, Value]) Range() Range[Key, Value] {
	return Range[Key, Value]{o, false, nil}
}

func (o *Range[Key, Value]) Next() bool {
	if !o.Started {
		o.Started = true
		o.Current = o.Map.List.Start
	} else {
		o.Current = o.Current.Next
	}
	return o.Current != nil
}

func (o *Range[TKey, TValue]) Key() (TKey, bool) {
	var zero TKey
	if o.Current == nil {
		return zero, false
	}
	return o.Current.Val, true
}

func (o *Range[TKey, TValue]) Val() (TValue, bool) {
	var zero TValue
	if o.Current == nil {
		return zero, false
	}
	return o.Map.Map[o.Current.Val].Val, true
}
