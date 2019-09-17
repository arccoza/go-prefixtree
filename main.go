package main

import (
	"fmt"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Hello World")
	t := New()
	t.Set("a1", 5)
	t.Set("a2", 4)
	t.Set("a3", 3)
	// t.Set("a11", 15)
	t.Set("a111", 115)
	t.Set("a112", 116)
	t.Set("a1121", 1116)
	t.Set("a1122", 1117)
	// spew.Dump(pm)
	// fmt.Println(pm)
	fmt.Println(t.Get("a").Values())

	for it := t.AsNode().Nodes(); it.Next(); {
		fmt.Println(it.Value().Value())
		// it.Close()
	}
}
