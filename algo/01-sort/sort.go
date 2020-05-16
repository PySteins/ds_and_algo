package _01_sort

import "github.com/gogf/gf/util/gutil"

type ISort interface {
	Sort([]interface{})
	compare(i1, i2 int) int
	swap(i1, i2 int)
}

type array []interface{}

func (a array) copy() (new []interface{}) {
	new = make([]interface{}, len(a))
	copy(new, a)
	return
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
