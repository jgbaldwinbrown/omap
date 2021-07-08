package main

import (
	"fmt"
	"local/jgbaldwinbrown/omap"
	"sort"
)

type ByNumKey struct {omap.Omapper}

func (b ByNumKey) Less(i, j int) bool {
	ki := b.Omapper.GetKey(i)
	kj := b.Omapper.GetKey(j)
	return ki.(int) < kj.(int)
}

func main() {
	o := omap.NewOmap()
	o.Set("a", 5)
	o.Set(22, "b")
	fmt.Println(o)
	for i:=0; i<o.Len(); i++ {
		val, ok := o.Get(o.GetKey(i))
		fmt.Println(i, o.GetKey(i), val, ok)
	}

	o.Insert(1, "new", "floop")
	fmt.Println(o)
	o.Swap(1,2)
	fmt.Println(o)
	o.Delete("a")
	fmt.Println(o)
	o.Delete(22)
	fmt.Println(o)
	o.Delete("new")
	fmt.Println(o)

	o2 := omap.NewOmap()
	b := ByNumKey{o2}
	b.Set(3, "three")
	b.Set(2, "twb")
	b.Set(1, "bne")
	fmt.Println(b.Omapper)
	sort.Sort(b)
	fmt.Println(b.Omapper)
}
