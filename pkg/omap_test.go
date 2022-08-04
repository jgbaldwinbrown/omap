package main

import (
	"testing"
	"fmt"
)

func TestMap(t *testing.T) {
	m := NewMap[string,float64]()
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
