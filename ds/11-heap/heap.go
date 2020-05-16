package _1_heap

import (
	"ds_and_algo/utils"
	"errors"

	"github.com/gogf/gf/util/gutil"
)

const DefaultCapacity = 10

type Heap struct {
	size       int
	elements   []interface{}
	comparator gutil.Comparator
}

func NewHeap(capacity int, comparator gutil.Comparator) *Heap {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}
	return &Heap{
		elements:   make([]interface{}, capacity),
		comparator: comparator,
	}
}

func Heapify(raw []interface{}, comparator gutil.Comparator) *Heap {
	h := &Heap{
		elements:   make([]interface{}, len(raw)),
		comparator: comparator,
	}
	// 最后一个非叶子节点开始下虑
	last := (h.size >> 1) - 1
	for i := last; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Clear() {
	for i := 0; i < h.size; i++ {
		h.elements[i] = nil
	}
}

func (h *Heap) Add(e interface{}) (err error) {
	if err = utils.ElementNotNilCheck(e); err != nil {
		return
	}
	h.ensureCapacity(h.size + 1)
	h.elements[h.size] = e
	h.size++
	h.siftUp(h.size - 1)
	return

}

func (h *Heap) Get() (interface{}, error) {
	if err := h.emptyCheck(); err != nil {
		return nil, err
	}
	return h.elements[0], nil
}

func (h *Heap) Remove() (first interface{}, err error) {
	if err := h.emptyCheck(); err != nil {
		return nil, err
	}
	h.size--
	first = h.elements[0]
	lastIndex := h.size
	h.elements[0] = h.elements[lastIndex]
	h.elements[lastIndex] = nil
	h.siftDown(0)
	return
}

func (h *Heap) Replace(new interface{}) (old interface{}) {

	old = h.elements[0]
	if old != nil {
		h.elements[0] = new
		h.siftDown(0)
	}
	return old

}

func (h *Heap) siftUp(index int) {
	e := h.elements[index]
	for index > 0 {
		parentIndex := (index - 1) >> 1
		p := h.elements[parentIndex]
		if h.comparator(e, p) <= 0 {
			break
		}
		h.elements[index] = p
		index = parentIndex
	}
	h.elements[index] = e
}

func (h *Heap) siftDown(index int) {
	e := h.elements[index]
	// 第一个叶子节点的索引 == 非叶子节点的数量
	half := h.size >> 1
	for index < half {
		leftIndex := (index << 1) + 1
		rightIndex := leftIndex + 1

		// 优先使用左节点
		childIndex := leftIndex
		child := h.elements[childIndex]
		if (rightIndex < h.size) && (h.comparator(h.elements[rightIndex], child) > 0) {
			// 右子节点存在且比较大
			childIndex = rightIndex
			child = h.elements[childIndex]
		}
		if h.comparator(e, child) >= 0 {
			break
		}

		h.elements[index] = child
		index = childIndex
	}
	h.elements[index] = e

}

func (h *Heap) ensureCapacity(capacity int) {
	oldCap := len(h.elements)
	if oldCap >= capacity {
		return
	}

	newCap := oldCap + (oldCap >> 1)
	newElements := make([]interface{}, newCap)
	for i := 0; i < h.size; i++ {
		newElements[i] = h.elements[i]
	}
	h.elements = newElements
}

func (h *Heap) emptyCheck() error {
	if h.IsEmpty() {
		return errors.New("heap is empty")
	}
	return nil
}
