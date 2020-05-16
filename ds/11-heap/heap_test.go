package _1_heap

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func Test(t *testing.T) {
	h := NewHeap(10, gutil.ComparatorInt)
	for i := 0; i < 1000; i++ {
		err := h.Add(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(h.Remove())
	fmt.Println(h.elements)
	fmt.Println(h.Remove())
	fmt.Println(h.elements)
	fmt.Println(h.Remove())
	fmt.Println(h.elements)
	fmt.Println(h.Replace(999))
	fmt.Println(h.elements)
	fmt.Println(h.Size())
}
