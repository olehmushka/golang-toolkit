package list

import (
	"testing"
)

func TestNewFIFOUniqueList(t *testing.T) {
	elements := []int{0, 1, 2, 3, 4, 5}
	list := NewFIFOUniqueList(10, func(left, right int) bool { return left == right }, elements...)
	if list.FirstElem.Value != elements[len(elements)-1] {
		t.Errorf("unexpected first element (actual=%d)", elements[len(elements)-1])
		return
	}
	if list.LastElem.Value != elements[0] {
		t.Errorf("unexpected last element (actual=%d)", elements[0])
		return
	}
	if list.Size != len(elements) {
		t.Errorf("unexpected list size (expected=%d, actual=%d)", len(elements), list.Size)
	}
}

func TestPush(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	if fe := listSize - 1; list.FirstElem.Value != fe {
		t.Errorf("unexpected first element (actual=%d)", fe)
		return
	}
	if le := 0; list.LastElem.Value != le {
		t.Errorf("unexpected last element (actual=%d)", le)
		return
	}
	if list.Size != listSize {
		t.Errorf("unexpected list size (expected=%d, actual=%d)", listSize, list.Size)
	}

	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	if fe := listSize - 1; list.FirstElem.Value != fe {
		t.Errorf("unexpected first element after second iter (actual=%d)", fe)
		return
	}
	if le := 0; list.LastElem.Value != le {
		t.Errorf("unexpected last element after second iter (actual=%d)", le)
		return
	}
	if list.Size != listSize {
		t.Errorf("unexpected list size after second iter (expected=%d, actual=%d)", listSize, list.Size)
	}

	shift := 10
	for i := 0; i < (listSize + shift); i++ {
		list.Push(i)
	}
	if fe := (listSize + shift) - 1; list.FirstElem.Value != fe {
		t.Errorf("unexpected first element after third iter (actual=%d)", fe)
		return
	}
	if le := shift; list.LastElem.Value != le {
		t.Errorf("unexpected last element after third iter (actual=%d)", le)
		return
	}
	if list.Size != listSize {
		t.Errorf("unexpected list size after third iter (expected=%d, actual=%d)", listSize, list.Size)
	}
}

func TestPop(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}
	list.Pop(listSize - 1)
	if newFE := listSize - 2; list.FirstElem.Value != newFE {
		t.Errorf("unexpected first element (expected=%d, actual=%d)", newFE, list.FirstElem.Value)
	}
	if newSize := listSize - 1; list.Size != newSize {
		t.Errorf("unexpected list size (expected=%d, actual=%d)", newSize, list.Size)
	}

	list.Pop(0)
	if newLE := 1; list.LastElem.Value != newLE {
		t.Errorf("unexpected last element (expected=%d, actual=%d)", newLE, list.LastElem.Value)
	}
	if newSize := listSize - 2; list.Size != newSize {
		t.Errorf("unexpected list size (expected=%d, actual=%d)", newSize, list.Size)
	}
}

func TestPopLast(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	list.PopLast()
	if newLE := 1; list.LastElem.Value != newLE {
		t.Errorf("unexpected last element (expected=%d, actual=%d)", newLE, list.LastElem.Value)
	}
	if newSize := listSize - 1; list.Size != newSize {
		t.Errorf("unexpected list size (expected=%d, actual=%d)", newSize, list.Size)
	}
}

func TestFindOne(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	actual := list.FindOne(func(prev, curr, next int) bool {
		return curr%2 == 0 && curr > 5
	})
	if expected := 6; actual == expected {
		t.Errorf("unexpected find one result (expected=%d, actual=%d)", expected, actual)
	}
}

func TestFilter(t *testing.T) {
	listSize := 10
	list := NewFIFOUniqueList(listSize, func(left, right int) bool { return left == right })
	for i := 0; i < listSize; i++ {
		list.Push(i)
	}

	actual := list.Filter(func(prev, curr, next int) bool {
		return curr%2 == 0
	})
	expected := []int{8, 6, 4, 2, 0}
	if len(expected) != len(actual) {
		t.Errorf("unexpected filter result length (expected=%d, actual=%d)", len(expected), len(actual))
		return
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("unexpected filter result[%d] (expected=%d, actual=%d)", i, expected[i], actual[i])
			return
		}
	}
}
