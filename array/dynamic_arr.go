package array

import (
	"errors"
)

const OutOfRangeByUpdate = "index error, the index must < array length and > 0"
const OutOfRangeByRead = "index error, the index must <= array length and > 0"

type Array struct {
	size int
	Data []int
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) Add(index int, value int) error {
	if index < 0 || index > a.size {
		return errors.New(OutOfRangeByUpdate)
	}
	for i := a.size - 1; i >= index; i-- {
		a.Data[i + 1] = a.Data[i]
	}
	a.Data[index] = value
	a.size++
	return nil
}

func (a *Array) AddFirst(value int) error {
	return a.Add(0, value)
}

func (a * Array) AddLast(value int) error {
	return a.Add(a.size-1, value)
}

func (a *Array) Set(index int, value int) error {
	if index < 0 || index >= a.size {
		return errors.New(OutOfRangeByRead)
	}
	a.Data[index] = value
	return nil
}

func (a *Array) Get(index int) (err error, res int) {
	if index < 0 || index >= a.size {
		return errors.New(OutOfRangeByRead), res
	}
	return err, a.Data[index]
}

func (a *Array) Remove(index int) (err error, res int) {
	if index < 0 || index >= a.size {
		return errors.New(OutOfRangeByRead), res
	}
	res = a.Data[index]
	for i := a.size - 1; i >= index; i-- {
		a.Data[i-1] = a.Data[i]
	}
	a.size--
	return nil, res
}

func (a *Array) RemoveElement(value int) error {
	_, index := a.Find(value)
	if index != -1 {
		_, _ = a.Remove(index)
	}
	return nil
}

func (a *Array) RemoveFirst() (err error, res int) {
	return a.Remove(0)
}

func (a *Array) RemoveLast() (err error, res int) {
	return a.Remove(a.size-1)
}

func (a Array) Find(value int) (err error, res int) {
	for i := 0; i <= a.size; i++ {
		if value == a.Data[i] {
			return nil, i
		}
	}
	return errors.New("not find"), -1
}

func NewArray(n int) *Array {
	data := make([]int, n)
	arr := &Array{
		size: 0,
		Data: data,
	}
	return arr
}