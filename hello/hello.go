package main

import (
	"fmt"
	"github.com/garudareiga/gocode/collection"
	"github.com/garudareiga/gocode/stringutil"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	var s collection.Stack
	s.Push("apple")
	s.Push("orange")

	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
}
