package _01_sort

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func TestBubbleSort(t *testing.T) {
	var l = array{
		55, 10, 23, 11, 45, 46, 22,
	}
	bs1 := NewBubbleSort(gutil.ComparatorInt, 1)
	bs1.Sort(l.copy())
	fmt.Println("======bs1=====")
	fmt.Println(bs1.array)
	fmt.Println(bs1.cmpCount)
	fmt.Println(bs1.swapCount)

	bs2 := NewBubbleSort(gutil.ComparatorInt, 2)
	bs2.Sort(l.copy())
	fmt.Println("======bs2=====")
	fmt.Println(bs2.array)
	fmt.Println(bs2.cmpCount)
	fmt.Println(bs2.swapCount)

	bs3 := NewBubbleSort(gutil.ComparatorInt, 3)
	bs3.Sort(l.copy())
	fmt.Println("======bs3=====")
	fmt.Println(bs3.array)
	fmt.Println(bs3.cmpCount)
	fmt.Println(bs3.swapCount)
}
