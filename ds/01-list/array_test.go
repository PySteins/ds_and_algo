package main

import (
	"testing"
)

func TestArrayInt_Print(t *testing.T) {
	l := NewArrayInt(10)
	var err error
	err = l.Add(1)
	if err != nil {
		t.Error(err)
	}
	err = l.Add(2)
	if err != nil {
		t.Error(err)
	}
	err = l.Add(0, 0)
	if err != nil {
		t.Error(err)
	}
	err = l.Add(3, 3)
	if err != nil {
		t.Error(err)
	}
	err = l.Add(4)
	if err != nil {
		t.Error(err)
	}
	err = l.Remove(0)
	if err != nil {
		t.Error(err)
	}
	l.Print()
}
