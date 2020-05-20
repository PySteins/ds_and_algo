package _01_sort

import (
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gutil"
)

type QuickSort struct {
	*sort
}

func NewQuickSort(comparator gutil.Comparator) *QuickSort {
	return &QuickSort{
		sort: NewSort(comparator),
	}
}

func (q *QuickSort) Sort(raw []interface{}) {
	if !q.check(raw) {
		return
	}
	q.quickSort(0, len(raw))
}

func (q *QuickSort) quickSort(start, end int) {
	if end-start < 2 {
		return
	}

	pivotIndex := q.getPivotIndex(start, end)
	q.quickSort(start, pivotIndex)
	q.quickSort(pivotIndex+1, end)
}

// 获取[start, end)范围内的轴点索引
func (q *QuickSort) getPivotIndex(start, end int) int {
	end--
	q.swap(start, grand.N(start, end))
	pivot := q.array[start]
	for start < end {

		// 先从右往左扫描
		for start < end {
			if q.compareElement(q.array[end], pivot) > 0 {
				end--
			} else {
				q.array[start] = q.array[end]
				start++
				// 换方向
				break
			}
		}

		// 从左往右
		for start < end {
			if q.compareElement(q.array[start], pivot) < 0 {
				start++
			} else {
				q.array[end] = q.array[start]
				end--
				// 换方向
				break
			}
		}
	}
	q.array[start] = pivot
	return start
}
