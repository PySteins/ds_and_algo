package _01_sort

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func TestHeapSort(t *testing.T) {
	var l = array{
		55, 10, 23, 11, 45, 46, 22,
	}
	hs := NewHeapSort(gutil.ComparatorInt)
	hs.Sort(l.copy())
	fmt.Println("======bs1=====")
	fmt.Println(hs.array)
	fmt.Println(hs.cmpCount)
	fmt.Println(hs.swapCount)
}
