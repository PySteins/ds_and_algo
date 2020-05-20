package _01_sort

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func TestGetTestSample(t *testing.T) {
	sample, sorted := GetTestSample(0, 10000, 32)
	sampleCopy := make([]interface{}, len(sample))
	var name string

	// name = "bubble_sort_1"
	// copy(sampleCopy, sample)
	// DoSort(name, NewBubbleSort(gutil.ComparatorInt, 1), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "bubble_sort_2"
	// copy(sampleCopy, sample)
	// DoSort(name, NewBubbleSort(gutil.ComparatorInt, 2), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "bubble_sort_3"
	// copy(sampleCopy, sample)
	// DoSort(name, NewBubbleSort(gutil.ComparatorInt, 3), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "selection_sort"
	// copy(sampleCopy, sample)
	// DoSort(name, NewSelectionSort(gutil.ComparatorInt), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "insertion_sort_1"
	// copy(sampleCopy, sample)
	// DoSort(name, NewInsertionSort(gutil.ComparatorInt, 1), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "insertion_sort_2"
	// copy(sampleCopy, sample)
	// DoSort(name, NewInsertionSort(gutil.ComparatorInt, 2), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	// name = "insertion_sort_3"
	// copy(sampleCopy, sample)
	// DoSort(name, NewInsertionSort(gutil.ComparatorInt, 3), sampleCopy)
	// if !reflect.DeepEqual(sampleCopy, sorted) {
	// 	fmt.Println(name, "测试未通过")
	// }

	name = "merge_sort"
	copy(sampleCopy, sample)
	DoSort(name, NewMergeSort(gutil.ComparatorInt), sampleCopy)
	if !reflect.DeepEqual(sampleCopy, sorted) {
		fmt.Println(name, "测试未通过")
	}

	name = "heap_sort"
	copy(sampleCopy, sample)
	DoSort(name, NewHeapSort(gutil.ComparatorInt), sampleCopy)
	if !reflect.DeepEqual(sampleCopy, sorted) {
		fmt.Println(name, "测试未通过")
	}

	name = "quick_sort"
	copy(sampleCopy, sample)
	DoSort(name, NewQuickSort(gutil.ComparatorInt), sampleCopy)
	if !reflect.DeepEqual(sampleCopy, sorted) {
		fmt.Println(name, "测试未通过")
	}

}
