package _01_sort

import (
	"github.com/gogf/gf/util/gutil"
)

type InsertionSort struct {
	*sort
	ver int
}

func NewInsertionSort(comparator gutil.Comparator, ver int) *InsertionSort {
	return &InsertionSort{
		NewSort(comparator),
		ver,
	}
}

func (i *InsertionSort) Sort(raw []interface{}) {
	if !i.check(raw) {
		return
	}
	switch i.ver {
	case 1:
		for begin := 1; begin < len(i.array); begin++ {
			currentIndex := begin
			for currentIndex > 0 && i.compare(currentIndex, currentIndex-1) < 0 {
				i.swap(currentIndex, currentIndex-1)
				currentIndex--
			}
		}
	case 2:
		// 将交换元素转为挪动元素 提高效率
		for begin := 1; begin < len(raw); begin++ {
			currentIndex := begin
			current := i.array[currentIndex]
			// 需要直接比较元素而非通过索引获取元素进行对比
			for currentIndex > 0 && i.compareElement(current, i.array[currentIndex-1]) < 0 {
				i.array[currentIndex] = i.array[currentIndex-1]
				currentIndex--
			}
			i.array[currentIndex] = current
		}
	case 3:
		// 将交换元素转为挪动元素 提高效率
		for begin := 1; begin < len(raw); begin++ {
			current := i.array[begin]
			// 对前部已排序的元素进行二分查找返回当前需要插入元素的位置
			// 减少比较次数
			// 插入位置后的排序元素往后移动一位
			insertIndex := i.BinarySearch(begin)
			for j := begin; j > insertIndex; j-- {
				i.array[j] = i.array[j-1]
			}
			i.array[insertIndex] = current
		}
	}
}

func (i *InsertionSort) compareElement(v1, v2 interface{}) int {
	i.cmpCount++
	return i.comparator(v1, v2)
}

func (i *InsertionSort) BinarySearch(endIndex int) int {
	begin := 0
	end := endIndex
	for end > begin {
		mid := (end + begin) >> 1
		sub := i.compareElement(i.array[mid], i.array[endIndex])
		switch {
		case sub > 0:
			end = mid
		case sub <= 0:
			begin = mid + 1
		}
	}
	return end
}
