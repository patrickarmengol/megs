// min-heap / priority queue
package binaryheap

import "fmt"

type Heap[T any] struct {
	data []T
	cmp  func(a T, b T) int
}

// New creates a new binary heap based on the passed cmp function.
//
// cmp should return a negative number when a < b, a positive number when a > b, and zero when a == b.
//
// This results in a min-heap. cmp can be wrapped with binaryheap.Reverse to make a max-heap instead.
func New[T any](cmp func(a T, b T) int) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		cmp:  cmp,
	}
}

// Of creates a new binary heap based on the passed cmp function, with the initial data.
//
// cmp should return a negative number when a < b, a positive number when a > b, and zero when a == b.
//
// This results in a min-heap. cmp can be wrapped with binaryheap.Reverse to make a max-heap instead.
func Of[T any](cmp func(a T, b T) int, vals ...T) *Heap[T] {
	heapify(vals, cmp)
	return &Heap[T]{
		data: vals,
		cmp:  cmp,
	}
}

func Reverse[T any](cmp func(a T, b T) int) func(a T, b T) int {
	return func(a T, b T) int {
		return cmp(b, a)
	}
}

func (h *Heap[T]) Push(val T) {
	h.data = append(h.data, val)
	percolateUp(h.data, len(h.data)-1, h.cmp)
}

func (h *Heap[T]) Pop() T {
	if h.IsEmpty() {
		panic("can't pop from empty heap")
	}

	r := h.data[0]

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	percolateDown(h.data, 0, h.cmp)

	return r
}

func (h *Heap[T]) Peek() T {
	if h.IsEmpty() {
		panic("can't peek on empty heap")
	}

	return h.data[0]
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *Heap[T]) String() string {
	return fmt.Sprintf("heap%v", h.data)
}

func percolateUp[T any](d []T, i int, cmp func(a T, b T) int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || cmp(d[i], d[parent]) > 0 {
			break
		}

		d[i], d[parent] = d[parent], d[i]
		i = parent
	}
}

func percolateDown[T any](d []T, i int, cmp func(a T, b T) int) {
	for {
		left, right := 2*i+1, 2*i+2
		if left >= len(d) || left < 0 {
			break
		}

		smaller := left
		if right < len(d) && cmp(d[left], d[right]) > 0 {
			smaller = right
		}

		if cmp(d[i], d[smaller]) <= 0 {
			break
		}

		d[i], d[smaller] = d[smaller], d[i]
		i = smaller
	}
}

func heapify[T any](d []T, cmp func(a T, b T) int) {
	for i := len(d)/2 - 1; i >= 0; i-- {
		percolateDown(d, i, cmp)
	}
}
