package collection

import (
	collection "." // use the package from the current directory
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	var s collection.Stack
	s.Push("apple")
	s.Push("orange")

	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
}
