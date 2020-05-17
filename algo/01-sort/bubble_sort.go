package _01_sort

import "github.com/gogf/gf/util/gutil"

type BubbleSort struct {
	*sort
	ver int
}

func NewBubbleSort(comparator gutil.Comparator, ver int) *BubbleSort {
	return &BubbleSort{
		NewSort(comparator),
		ver,
	}
}

func (b *BubbleSort) Sort(raw []interface{}) {
	if !b.check(raw) {
		return
	}
	// 确保方法完整性，不提取重复代码
	switch b.ver {
	case 1:
		// 默认
		l := len(b.array) - 1
		for l > 0 {
			for begin := 1; begin <= l; begin++ {
				if b.compare(begin-1, begin) > 0 {
					b.swap(begin-1, begin)
				}
			}
			l--
		}
	case 2:
		// 完全有序可提前退出
		l := len(b.array) - 1
		for l > 0 {
			sorted := true
			for begin := 1; begin <= l; begin++ {
				if b.compare(begin-1, begin) > 0 {
					b.swap(begin-1, begin)
					sorted = false
				}
			}
			if sorted {
				break
			}
			l--
		}
	case 3:
		// 尾部局部有序记录交换位置，减少比较次数
		l := len(b.array) - 1
		for l > 0 {
			var sortedIndex int
			for begin := 1; begin <= l; begin++ {
				if b.compare(begin-1, begin) > 0 {
					b.swap(begin-1, begin)
					sortedIndex = begin
				}
			}
			l = sortedIndex
		}
	}
}
