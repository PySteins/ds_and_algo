package _01_sort

import "github.com/gogf/gf/util/gutil"

type MergeSort struct {
	*sort
	leftArray []interface{}
}

func NewMergeSort(comparator gutil.Comparator) *MergeSort {
	return &MergeSort{
		sort: NewSort(comparator),
	}
}

func (m *MergeSort) Sort(raw []interface{}) {
	if !m.check(raw) {
		return
	}
	m.leftArray = make([]interface{}, len(raw)>>1)
	m.divide(0, len(raw))
}

func (m *MergeSort) divide(begin, end int) {
	if end-begin < 2 {
		return
	}
	mid := (begin + end) >> 1
	m.divide(begin, mid)
	m.divide(mid, end)
	m.merge(begin, mid, end)
}

func (m *MergeSort) merge(begin, mid, end int) {
	leftIndex, leftEnd := 0, mid-begin
	rightIndex, rightEnd := mid, end
	arrayIndex := begin

	for i := 0; i < leftEnd; i++ {
		m.leftArray[i] = m.array[begin+i]
	}

	for leftIndex < leftEnd {
		// 同时需要确保稳定性
		if rightIndex < rightEnd && m.compareElement(m.array[rightIndex], m.leftArray[leftIndex]) < 0 {
			m.array[arrayIndex] = m.array[rightIndex]
			rightIndex++
		} else {
			m.array[arrayIndex] = m.leftArray[leftIndex]
			leftIndex++
		}
		arrayIndex++
	}
}
