package _01_sort

import "github.com/gogf/gf/util/gutil"

type SelectionSort struct {
	*sort
}

func NewSelectionSort(comparator gutil.Comparator) *SelectionSort {
	return &SelectionSort{
		NewSort(comparator),
	}
}

func (s *SelectionSort) Sort(raw []interface{}) {
	if !s.check(raw) {
		return
	}
	l := len(s.array) - 1
	for l > 0 {
		var maxIndex int
		for i := 1; i <= l; i++ {
			if s.compare(maxIndex, i) <= 0 {
				maxIndex = i
			}
		}
		s.swap(maxIndex, l)
		l--
	}
}
