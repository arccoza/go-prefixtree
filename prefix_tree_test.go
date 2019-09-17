package main

import (
	// "os"
	// "flag"
	// "testing"
	"fmt"
)

// func TestMain(m *testing.M) {
// 	flag.Parse()
// 	os.Exit(m.Run())
// }

func ExamplePrefixTree() {
	t := New()
	t.Set("/root/home/bob/test.txt", "test.txt")
	t.Set("/root/home/bob/test2.txt", "test2.txt")
	t.Set("/root/home/bob/Videos/Futurama/S01E01.m4v", "S01E01.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S01E02.m4v", "S01E02.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S02E01.m4v", "S02E01.m4v")
	t.Set("/root/home/bob/Videos/Futurama/S02E02.m4v", "S02E02.m4v")

	fmt.Println(t.Get("/root/").Value())
	fmt.Println(t.Get("/root/").Values())
	fmt.Println(t.Get("/root/home/bob/Videos/Futurama/S02E02.m4v").Value())
	fmt.Println(t.Get("/root/home/bob/Videos/Futurama/").Values())
	// Output:
	// <nil>
	// [test.txt test2.txt S02E01.m4v S02E02.m4v S01E01.m4v S01E02.m4v]
	// S02E02.m4v
	// [S01E01.m4v S01E02.m4v S02E01.m4v S02E02.m4v]
}
