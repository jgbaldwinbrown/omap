package omap

type Omapper interface {
	Get(interface{}) (interface{}, error)
	GetKey(int) interface{}
	Set(interface{}, interface{})
	Len() int
	Delete(interface{})
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
	}
}