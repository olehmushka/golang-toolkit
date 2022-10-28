package list

type Element[T any] struct {
	Prev  *Element[T]
	Value T
	Next  *Element[T]
}

type FIFOUniqueList[T any] struct {
	FirstElem   *Element[T]
	LastElem    *Element[T]
	Size        int
	MaxSize     int
	CompareFunc func(T, T) bool
}

func NewFIFOUniqueList[T any](maxSize int, compareFnc func(T, T) bool, elements ...T) *FIFOUniqueList[T] {
	out := &FIFOUniqueList[T]{
		MaxSize:     maxSize,
		CompareFunc: compareFnc,
	}
	for i := 0; i < len(elements); i++ {
		out.Push(elements[i])
	}

	return out
}

func (l *FIFOUniqueList[T]) Push(in T) {
	if l.FirstElem == nil {
		el := &Element[T]{
			Value: in,
		}
		l.FirstElem = el
		l.LastElem = el
		l.Size++

		return
	}
	_ = l.Pop(in)

	el := &Element[T]{
		Value: in,
		Next:  l.FirstElem,
	}
	l.FirstElem.Prev = el
	l.FirstElem = el
	l.Size++

	if l.MaxSize < l.Size {
		l.PopLast()
	}
}

func (l *FIFOUniqueList[T]) Pop(in T) bool {
	for el := l.FirstElem; el != nil; el = el.Next {
		if l.CompareFunc(el.Value, in) {
			if el.Prev != nil && el.Next != nil {
				el.Next.Prev = el.Prev
				el.Prev.Next = el.Next
				l.Size--

				return true
			}
			if el.Prev == nil && el.Next != nil {
				el.Next.Prev = nil
				l.FirstElem = el.Next
				l.Size--

				return true
			}
			if el.Prev != nil && el.Next == nil {
				el.Prev.Next = nil
				l.LastElem = el.Prev
				l.Size--

				return true
			}
		}
	}

	return false
}

func (l *FIFOUniqueList[T]) PopLast() {
	if l.Size == 0 {
		return
	}
	if l.Size == 1 {
		l.FirstElem = nil
		l.LastElem = nil
		l.Size = 0

		return
	}
	l.LastElem = l.LastElem.Prev
	l.LastElem.Next = nil
	l.Size--
}

func (l *FIFOUniqueList[T]) FindOne(cb func(prev, curr, next T) bool) T {
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			return el.Value
		}
	}
	var zero T

	return zero
}

func (l *FIFOUniqueList[T]) Filter(cb func(prev, curr, next T) bool) []T {
	out := make([]T, 0, l.Size)
	for el := l.FirstElem; el != nil; el = el.Next {
		var prev, next T
		if el.Prev != nil {
			prev = el.Prev.Value
		}
		if el.Next != nil {
			next = el.Next.Value
		}
		if cb(prev, el.Value, next) {
			out = append(out, el.Value)
		}
	}

	return out
}
