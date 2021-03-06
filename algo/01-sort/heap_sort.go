package _01_sort

import (
	"github.com/gogf/gf/util/gutil"
)

type HeapSort struct {
	*sort
	heapSize int
}

func NewHeapSort(comparator gutil.Comparator) *HeapSort {
	return &HeapSort{
		sort: NewSort(comparator),
	}
}

func (h *HeapSort) Sort(raw []interface{}) {
	if !h.check(raw) {
		return
	}
	h.heapSize = len(raw)
	// 最后一个非叶子节点开始下虑
	last := (h.heapSize >> 1) - 1
	for begin := last; begin >= 0; begin-- {
		h.siftDown(begin)
	}
	for h.heapSize > 1 {
		h.swap(0, h.heapSize-1)
		h.heapSize--
		h.siftDown(0)
	}

}

func (h *HeapSort) siftDown(index int) {
	e := h.array[index]
	// 第一个叶子节点的索引 == 非叶子节点的数量
	half := h.heapSize >> 1
	for index < half {
		leftIndex := (index << 1) + 1
		rightIndex := leftIndex + 1

		// 优先使用左节点
		childIndex := leftIndex
		child := h.array[childIndex]
		if (rightIndex < h.heapSize) && (h.compareElement(h.array[rightIndex], child) > 0) {
			// 右子节点存在且比较大
			childIndex = rightIndex
			child = h.array[childIndex]
		}
		if h.compareElement(e, child) >= 0 {
			break
		}

		h.array[index] = child
		index = childIndex
	}
	h.array[index] = e

}
