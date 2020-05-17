package _01_sort

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func TestInsertionSort(t *testing.T) {
	var l = array{
		55, 10, 23, 11, 45, 46, 22,
	}
	is := NewInsertionSort(gutil.ComparatorInt, 3)
	is.Sort(l.copy())
	fmt.Println("======bs1=====")
	fmt.Println(is.array)
	fmt.Println(is.cmpCount)
	fmt.Println(is.swapCount)
}
