package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ArrayInt struct {
	*AbstractList
	Elements []interface{}
}

func NewArrayInt(capacity int) *ArrayInt {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}
	return &ArrayInt{
		AbstractList: NewAbstractList(capacity),
		Elements:     make([]interface{}, capacity),
	}
}

func (a *ArrayInt) Len() int {
	return a.Size
}

func (a *ArrayInt) Add(v interface{}, index ...int) error {
	var idx int
	if len(index) > 0 {
		idx = index[0]
		if err := a.IndexCheckForAdd(idx); err != nil {
			return err
		}
		for i := a.Size - idx; i > idx; i-- {
			a.Elements[i] = a.Elements[i-1]
		}
	} else {
		idx = a.Size
	}
	a.Elements[idx] = v
	a.Size++
	return nil
}

func (a *ArrayInt) Remove(index int) error {
	if err := a.IndexCheck(index); err != nil {
		return err
	}
	for i := index; i < a.Size; i++ {
		a.Elements[i] = a.Elements[i+1]
	}
	a.Size--
	return nil
}

func (a *ArrayInt) Set(index int, v interface{}) error {
	if err := a.IndexCheck(index); err != nil {
		return err
	}
	a.Elements[index] = v
	return nil
}

func (a *ArrayInt) Get(index int) (err error, v interface{}) {
	if err := a.IndexCheck(index); err != nil {
		return err, v
	}
	v = a.Elements[index]
	return
}

func (a *ArrayInt) Print() {
	str := strings.Builder{}
	str.WriteString("[")
	for i := 0; i < a.Size; i++ {
		if i != 0 {
			str.WriteString(", ")
		}
		v := a.Elements[i]
		switch v.(type) {
		case int:
			str.WriteString(strconv.Itoa(v.(int)))
		case string:
			str.WriteString(v.(string))
		default:
			str.WriteString("unknown_type")
		}
	}
	str.WriteString("]")
	fmt.Println(str.String())
}
