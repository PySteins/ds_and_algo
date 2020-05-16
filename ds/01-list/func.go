package main

import "errors"

var (
	DefaultCapacity = 10
	OutOfIndexErr   = errors.New("数组越界")
)

type List interface {
	Len() int
	Add(interface{}, ...int) error
	Remove(int) error
	Set(int, interface{}) error
	Get(int) (error, interface{})
	Print()
}

type AbstractList struct {
	Capacity int
	Size     int
}

func NewAbstractList(capacity int) *AbstractList {
	return &AbstractList{Capacity: capacity, Size: 0}
}

func (a *AbstractList) IndexCheck(index int) error {
	if index >= a.Size {
		return OutOfIndexErr
	}
	return nil
}

func (a *AbstractList) IndexCheckForAdd(index int) error {
	if index > a.Size {
		return OutOfIndexErr
	}
	return nil
}
