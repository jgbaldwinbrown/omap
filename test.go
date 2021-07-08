package main

import (
	"fmt"
	"local/jgbaldwinbrown/omap"
	// "sort"
)

// type OmapSortable *omap.Omapper
// 
// func (o OmapSortable) Swap(i1 int, i2 int) {
// 	k1 = o.GetKey(i1)
// 	k2 = o.GetKey(i2)
// 	v1, _ = o.Get(k1)
// 	v2, _ = o.Get(k2)
// 	o.Delete(k1)
// 	o.Delete(k2)
// 	o.Insert(i1, k2, v2)
// 	o.
// }
// 
// type ByNum struct {omap.Omapper}
// 
// func (b ByNum) Less(i, j int) bool {return b.Omapper.Get(i) < b.Omapper.Get(j)}

func main() {
	o := omap.NewOmap()
	o.Set("a", 5)
	o.Set(22, "b")
	fmt.Println(o)
	for i:=0; i<o.Len(); i++ {
		val, ok := o.Get(o.GetKey(i))
		fmt.Println(i, o.GetKey(i), val, ok)
	}
	o.Delete("a")
	fmt.Println(o)
	o.Delete(22)
	fmt.Println(o)

	// o2 := OmapSortable()
	// o2.Set(3, "three")
	// o2.Set(2, "two2")
	// o2.Set(1, "o2ne")
	// fmt.Println(o2)
	// sort.Sort(o2)
	// fmt.Println(o2)
}
