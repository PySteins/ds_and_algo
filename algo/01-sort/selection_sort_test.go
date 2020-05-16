package _01_sort

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func Test(t *testing.T) {
	var l = array{
		55, 10, 23, 11, 45, 46, 22,
	}
	ss := NewSelectionSort(gutil.ComparatorInt)
	ss.Sort(l.copy())
	fmt.Println("======bs1=====")
	fmt.Println(ss.array)
	fmt.Println(ss.cmpCount)
	fmt.Println(ss.swapCount)

	i := make([]interface{}, 10)
	fmt.Println(i)
}
