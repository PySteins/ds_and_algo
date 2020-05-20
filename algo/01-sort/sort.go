package _01_sort

import (
	"fmt"
	s "sort"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gutil"
)

type ISort interface {
	Sort([]interface{})
	compare(i1, i2 int) int
	swap(i1, i2 int)
	compareElement(v1, v2 interface{}) int
	operation() (int, int)
}

type sort struct {
	array      []interface{}
	cmpCount   int
	swapCount  int
	comparator gutil.Comparator
}

func NewSort(comparator gutil.Comparator) *sort {
	return &sort{comparator: comparator}
}

func (s *sort) check(raw []interface{}) bool {
	s.array = raw
	if raw == nil || len(raw) < 2 {
		return false
	}
	return true
}

func (s *sort) Sort([]interface{}) {}

func (s *sort) compare(i1, i2 int) int {
	s.cmpCount++
	return s.comparator(s.array[i1], s.array[i2])
}

func (s *sort) swap(i1, i2 int) {
	s.swapCount++
	tmp := s.array[i1]
	s.array[i1] = s.array[i2]
	s.array[i2] = tmp
}

func (s *sort) compareElement(v1, v2 interface{}) int {
	s.cmpCount++
	return s.comparator(v1, v2)
}

func (s *sort) operation() (int, int) {
	return s.cmpCount, s.swapCount
}

func GetTestSample(min, max, num int) (array []interface{}, sorted []interface{}) {
	array = make([]interface{}, num)
	for i := 0; i < num; i++ {
		array[i] = grand.N(min, max)
	}
	sorted = make([]interface{}, num)
	copy(sorted, array)
	tmp := gconv.SliceInt(sorted)
	s.Ints(tmp)
	sorted = gconv.SliceAny(tmp)
	return
}

func DoSort(name string, sortType ISort, raw []interface{}) {
	fmt.Println(name, "===========")
	start := gtime.TimestampMilli()
	sortType.Sort(raw)
	fmt.Println("time spent:", gtime.TimestampMilli()-start, "ms")
	compare, swap := sortType.operation()
	fmt.Println("compare times:", compare, "次")
	fmt.Println("swap times:   ", swap, "次")
	fmt.Println("==================")

}
