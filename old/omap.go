package omap

type Omapper interface {
	Get(interface{}) (interface{}, bool)
	GetKey(int) interface{}
	Set(interface{}, interface{})
	Len() int
	Delete(interface{})
	Insert(int, interface{}, interface{})
	Swap(int, int)
}

type omapEntry struct {
	index int
	value interface{}
}

type Omap struct {
	keys []interface{}
	omap map[interface{}]omapEntry
}

func NewOmap() (out *Omap) {
	out = new(Omap)
	out.keys = make([]interface{}, 0)
	out.omap = make(map[interface{}]omapEntry)
	return out
}

func (o *Omap) Get(k interface{}) (interface{}, bool) {
	e, ok := o.omap[k]
	return e.value, ok
}

func (o *Omap) GetKey(i int) (v interface{}) {
	return o.keys[i]
}

func (o *Omap) Set(k interface{}, v interface{}) {
	old, ok := o.omap[k]
	var index int
	if ! ok {
		index = len(o.keys)
		o.keys = append(o.keys, k)
	} else {
		index = old.index
	}
	o.omap[k] = omapEntry{index: index, value: v}
}

func (o *Omap) Len() int {
	return len(o.omap)
}

func (o *Omap) Delete(k interface{}) {
	e, ok := o.omap[k]
	if ok {
		delete(o.omap, k)
		o.keys = append(o.keys[:e.index], o.keys[e.index+1:]...)
		for i:=e.index; i<len(o.keys); i++ {
			e, _ = o.omap[o.keys[i]]
			e.index--
			o.omap[o.keys[i]] = e
		}
	}
}

func (o *Omap) Insert(i int, k interface{}, v interface{}) {
	_, ok := o.omap[k]
	if ok {
		o.Delete(k)
	}
	o.keys = append(o.keys[:i+1], o.keys[i:]...)
	o.keys[i] = k
	o.omap[k] = omapEntry{i, v}
	for index := i+1; index < len(o.keys); index++ {
		new_k := o.keys[index]
		e, _ := o.omap[new_k]
		e.index++
		o.omap[new_k] = e
	}
}

func (o *Omap) Swap(i1 int, i2 int) {
	k1 := o.keys[i1]
	k2 := o.keys[i2]
	o.keys[i1], o.keys[i2] = k2, k1

	e, _ := o.omap[k1]
	e.index = i2
	o.omap[k1] = e

	e, _ = o.omap[k2]
	e.index = i1
	o.omap[k2] = e
}
