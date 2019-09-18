package main

import (
	// "os"
	// "flag"
	// "testing"
	"fmt"
	"sort"
)

// func TestMain(m *testing.M) {
// 	flag.Parse()
// 	os.Exit(m.Run())
// }

func sortValues(vs []interface{}) []interface{} {
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].(string) < vs[j].(string)
	})
	return vs
}

func ExamplePrefixTree() {
	t := New()
	t.Set("/root/home/bob/test.txt", "test.txt")
	t.Set("/root/home/bob/test2.txt", "test2.txt")
	t.Set("/root/home/bob/Videos/Futurama/S01E01.m4v", "S01E01.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S01E02.m4v", "S01E02.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S02E01.m4v", "S02E01.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S02E02.m4v", "S02E02.m4v")

	fmt.Println(t.Get("/root/").Value())
	fmt.Println(sortValues(t.Get("/root/").Values()))
	fmt.Println(t.Get("/root/home/bob/Videos/Futurama/S02E02.m4v").Value())
	fmt.Println(sortValues(t.Get("/root/home/bob/Videos/Futurama/").Values()))
	// Output:
	// <nil>
	// [S01E01.m4v S01E02.m4v S02E01.m4v S02E02.m4v test.txt test2.txt]
	// S02E02.m4v
	// [S01E01.m4v S01E02.m4v S02E01.m4v S02E02.m4v]
}

func ExamplePrefixTree_Set() {
	t := New()

	fmt.Println(t.Len())
	fmt.Println(t.Get("1234").Value())
	t.Set("1234", "some-value") // Set a value at a key
	fmt.Println(t.Len())
	fmt.Println(t.Get("1234").Value())
	fmt.Println(t.Get("1234").Values())
	t.Set("123456", "some-other-value")  // Set another value and a (deeper) key
	t.Set("123444", "yet-another-value") // Set another value on a different branch
	fmt.Println(t.Len())
	fmt.Println(t.Get("1234").Value())
	fmt.Println(t.Get("1234").Values())
	fmt.Println(t.Get("12345").Value())
	fmt.Println(t.Get("123456").Value())
	t.Set("1234", "updated-value") // Update the value at "1234"
	fmt.Println(t.Get("1234").Value())
	fmt.Println(t.Get("123456").Value())
	// Output:
	// 0
	// <nil>
	// 1
	// some-value
	// [some-value]
	// 2
	// some-value
	// [some-value some-other-value yet-another-value]
	// <nil>
	// some-other-value
	// updated-value
	// some-other-value
}

func ExamplePrefixTree_Get() {
	t := New()
	t.Set("1234", "some-value")

	fmt.Println(t.Get("1234"))
	fmt.Println(t.Get("1234").Value())
	fmt.Println(t.Get("12"))
	fmt.Println(t.Get("12").Value())
	fmt.Println(t.Get("12").Values())
	// Output:
	// Node{key: 4, parent: 3, value: some-value, isValue: true, children: 0, len: 1}
	// some-value
	// Node{key: 2, parent: 1, value: <nil>, isValue: false, children: 1, len: 1}
	// <nil>
	// [some-value]
}
